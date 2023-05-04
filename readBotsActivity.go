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


func readBotsActivity(username string, monitor string) ([]BotDocumentStruct, error) {
	
	client, err := Client()
	ctx, _ 		:= context.WithTimeout(context.Background(), 10*time.Second)
	err 	= client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	// filter := bson.D{
	// 	{Key: "lastupdate", Value: bson.D{
	// 		{Key: "$gte", Value: hour0},
	// 		{Key: "$lte", Value: hour24},
	// 	}},
	// }

	filter 	:= bson.D {
		{Key: "username", Value: username},
		{Key: "monitor", Value: monitor},
	}

	collection 			:= client.Database(DATABASE).Collection(COLLECTIONBOTACTIVITY)
	result, errorRead 	:= collection.Find(ctx, filter)
	if errorRead != nil {
		return nil, errorRead
	}
	defer result.Close(ctx)

	var res []BotDocumentStruct
	for result.Next(ctx) {
		var bot BotDocumentStruct
		err		:= result.Decode(&bot)
		if err 	!= nil {
			log.Fatal(err)
		}
		res = append(res, bot)
	}

	return res, err
}

func readBotsActivitybyStatus(username string, monitor string, status string) ([]BotDocumentStruct, error) {
	client, err := Client()
	ctx, _ 		:= context.WithTimeout(context.Background(), 10*time.Second)
	err 	= client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	filter 	:= bson.D {
		{Key: "username", Value: username},
		{Key: "monitor", Value: monitor},
		{Key: "status", Value: status},
	}

	collection 			:= client.Database(DATABASE).Collection(COLLECTIONBOTACTIVITY)
	result, errorRead 	:= collection.Find(ctx, filter)
	if errorRead != nil {
		return nil, errorRead
	}
	defer result.Close(ctx)

	var res []BotDocumentStruct
	for result.Next(ctx) {
		var bot BotDocumentStruct
		err		:= result.Decode(&bot)
		if err 	!= nil {
			log.Fatal(err)
		}
		res = append(res, bot)
	}

	return res, err
}


func findBotsbyStatus(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	queryUsername 	:= c.Query("username")
	queryMonitor	:= c.Query("monitor")
	queryPassword	:= c.Query("password")
	queryStatus		:= c.Query("status")

	if queryUsername == "" || queryMonitor == "" || queryPassword == "" || queryStatus == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	if BotStatus[queryStatus] == "" {
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
	if user.Password != queryPassword {
		c.Status(http.StatusUnauthorized)
		return
	}

	result, err := readBotsActivitybyStatus(queryUsername, queryMonitor, queryStatus)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		if err == mongo.ErrNoDocuments {
			
			c.IndentedJSON(http.StatusOK, &ResponseFindBot{})
		}
		return
	}

	Bots := ResponseFindBot{}
	Bots.List = result

	c.IndentedJSON(http.StatusOK, Bots)
}


func getBotsActivity(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	queryUsername 	:= c.Query("username")
	queryMonitor	:= c.Query("monitor")
	queryPassword	:= c.Query("password")
	// queryYear 		:= c.Query("y")
	// queryMonth 		:= c.Query("m")
	// queryDay 		:= c.Query("d")

	if queryUsername == "" || queryMonitor == "" || queryPassword == "" {
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
	if user.Password != queryPassword {
		c.Status(http.StatusUnauthorized)
		return
	}

	// year, errYear 	:= strconv.Atoi(queryYear)
	// month, errMonth := strconv.Atoi(queryMonth)
	// day, errDay 	:= strconv.Atoi(queryDay)

	// if errYear != nil || errMonth != nil || errDay != nil {
	// 	c.Status(http.StatusBadRequest)
	// 	return
	// }

	// timeFilter 	:= time.Date(year, time.Month(month), day, 0, 0, 1, 0, *&time.UTC)
	// hour0 		:= time.Date(timeFilter.Year(), timeFilter.Month(), timeFilter.Day(), 0, 0, 1, 0, *&time.UTC).UnixMilli()
	// hour24 		:= time.Date(timeFilter.Year(), timeFilter.Month(), timeFilter.Day(), 23, 59, 59, 0, *&time.UTC).UnixMilli()

	result, err := readBotsActivity(queryUsername, queryMonitor)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		if err == mongo.ErrNoDocuments {
			c.IndentedJSON(http.StatusOK, &ResponseFindBot{})
		}
		return
	}

	Bots := ResponseFindBot{}
	Bots.List = result

	c.IndentedJSON(http.StatusOK, Bots)
}