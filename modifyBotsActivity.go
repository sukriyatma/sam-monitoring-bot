package main

import (
	"context"
	"log"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func insertBotsActivity(data []interface{}) (*mongo.InsertManyResult, error) {
	client, err := Client()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err 	= client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)

	collection 			:= client.Database(DATABASE).Collection(COLLECTIONBOTACTIVITY)
	result, insertErr 	:= collection.InsertMany(ctx, data)

	return result, insertErr
}

func insertMonitor(username string ,monitor string) int64 {
	client, err := Client()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return 0
	}

	filter := bson.D{{Key: "username", Value : username}}
	update := bson.D{{Key: "$push", Value: bson.D{
		{Key: "monitors", Value: monitor},
	}}}

	collection := client.Database(DATABASE).Collection(COLLECTIONUSER)
	result, insertErr := collection.UpdateOne(ctx, filter, update)
	if insertErr != nil {
		log.Fatal(insertErr)
	}

	return result.ModifiedCount
}


func findUpdateBotActivity(bot BotStruct, monitor string, epochMilisecond int64) (*BotDocumentStruct, error){
	client, err := Client()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	timeFilter := time.UnixMilli(epochMilisecond)
	hour0 := time.Date(timeFilter.Year(), timeFilter.Month(), timeFilter.Day(), 0, 0, 1, 0, *&time.UTC).UnixMilli()
	hour24 := time.Date(timeFilter.Year(), timeFilter.Month(), timeFilter.Day(), 23, 59, 59, 0, *&time.UTC).UnixMilli()

	filter := bson.D{
		{Key: "name", Value: bot.Name},
		{Key: "monitor", Value: monitor},
		{Key: "lastupdate", Value: bson.D{
			{Key: "$gte", Value: hour0},
			{Key: "$lte", Value: hour24},
		} },
	}

	update := bson.D { 
		{Key: "$set", Value: bson.D {
			{Key: "lastupdate", Value: epochMilisecond},
			{Key: "captcha", Value: bot.Captcha},
			{Key: "level", Value: bot.Level},
			{Key: "status", Value: bot.Status},
			{Key: "world", Value: bot.World},
			{Key: "x", Value: bot.X},
			{Key: "y", Value: bot.Y},
			{Key: "profit", Value: bot.Profit},
		}}}

	var res *BotDocumentStruct
	collection := client.Database(DATABASE).Collection(COLLECTIONBOTACTIVITY)
	errorRead := collection.FindOneAndUpdate(ctx, filter, update).Decode(&res)
	
	return res, errorRead
}


// Main func
func postBotsActivity(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var Body BodyStruct
	if err := c.BindJSON(&Body); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	
	queryUsername := c.Query("username")
	if queryUsername == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	
	user, errFindUser := findUser(queryUsername)
	if errFindUser != nil {
		c.Status(http.StatusInternalServerError)
		if errFindUser == mongo.ErrNoDocuments {
			c.Status(http.StatusNotFound)
		}		
		return 
	}

	isMonitor := func () bool {
		for i:=0; i<len(user.Monitors);i++ {
			if user.Monitors[i] == Body.Monitor {
				return true
			}
		}
		return false
	}
	
	if user.Password != Body.Password {
		c.Status(http.StatusUnauthorized)
		return
	}
	
	data := []interface{} {}
	successInsertUpdate := 0
	epochMilisecond := time.Now().UnixMilli()
	listBot := Body.List
	for i:= 0; i < len(listBot); i++ {
		if BotStatus[listBot[i].Status] == "" {
			c.Status(http.StatusBadRequest)
			return
		}

		bot := &BotDocumentStruct{}
		bot.Username	= queryUsername
		bot.Monitor 	= Body.Monitor
		bot.LastUpdate 	= epochMilisecond
		bot.Name		= listBot[i].Name
		bot.Status		= listBot[i].Status
		bot.World		= listBot[i].World
		bot.X			= listBot[i].X
		bot.Y			= listBot[i].Y
		bot.Captcha 	= listBot[i].Captcha
		bot.Level		= listBot[i].Level
		bot.Profit		= listBot[i].Profit		

		_, errFind := findUpdateBotActivity(listBot[i], Body.Monitor, epochMilisecond)
		if errFind != nil && errFind == mongo.ErrNoDocuments{
			data = append(data, bot)
		}
		
		successInsertUpdate += 1
	}
	
	if !isMonitor()  {
		defer insertMonitor(queryUsername, Body.Monitor)
	}

	// Insert all document not ready yet
	if len(data) > 0 {
		res, errInsert := insertBotsActivity(data)
		if errInsert != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		successInsertUpdate += len(res.InsertedIDs)
	}
	
	c.IndentedJSON(http.StatusCreated, successInsertUpdate)	
}

