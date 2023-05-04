package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func findMonitors(c *gin.Context) {

	queryUsername := c.Query("username")
	queryPassword := c.Query("password")

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if queryUsername == "" || queryPassword == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	user, errFind := findUser(queryUsername)
	if errFind != nil {
		if errFind == mongo.ErrNoDocuments {
			c.Status(http.StatusNotFound)
			return
		}
	}

	if user.Password != queryPassword {
		c.Status(http.StatusUnauthorized)
		return
	}
	c.IndentedJSON(http.StatusOK, user.Monitors)
}

func deleteMonitor(username string, password string ,monitor string) ( *mongo.UpdateResult, error) {

	client, err := Client()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(DATABASE).Collection(COLLECTIONUSER)
	
	filter := bson.D{
		{Key: "username", Value: username},
		{Key: "password", Value: password},
	}
	update := bson.D{{Key: "$pull", Value: bson.D{
		{Key: "monitors", Value: monitor},
	}}} 	

	return collection.UpdateOne(ctx, filter, update) 
}

func removeMonitor(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	queryUsername := c.Query("username")
	queryPassword := c.Query("password")
	queryMonitor  := c.Query("monitor")

	if queryUsername == "" || queryPassword == "" || queryMonitor == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	user, errFind := findUser(queryUsername)
	if errFind != nil {
		if errFind == mongo.ErrNoDocuments {
			c.Status(http.StatusUnauthorized)
			return
		}
	}

	if user.Password != queryPassword {
		c.Status(http.StatusUnauthorized)
		return
	}

	_, error := deleteMonitor(queryUsername, queryPassword, queryMonitor)
	if error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusAccepted)

}
