package main

import (
	"context"
	
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func insertUser(data UserDocumentStruct) (*mongo.InsertOneResult, error) {
	client, err := Client()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err 	= client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)

	collection 			:= client.Database(DATABASE).Collection(COLLECTIONUSER)
	result, insertErr 	:= collection.InsertOne(ctx, data)

	return result, insertErr
}

func findUser(username string) (*UserDocumentStruct, error) {
	client, err := Client()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(DATABASE).Collection(COLLECTIONUSER)
	
	var user *UserDocumentStruct
	err = collection.FindOne(ctx, bson.D{{Key: "username", Value: username}}).Decode(&user)

	return user, err
}


func validateUser(c *gin.Context) {

	queryUsername := c.Query("username")
	queryPassword := c.Query("password")

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if queryUsername == "" || queryPassword == "" {
		// c.IndentedJSON(http.StatusBadRequest, 
			// HttpBadRequest{Code: http.StatusBadRequest, Status: "BadRequest"})
		c.Status(http.StatusBadRequest)
		return
	}

	user, errFind := findUser(queryUsername)
	if errFind != nil {
		if errFind == mongo.ErrNoDocuments {
			// c.IndentedJSON(http.StatusNotFound, 
				// HttpNotFound{Code: http.StatusNotFound, Status: "Not Found"})
			c.Status(http.StatusNotFound)
			return
		}
	}
	
	if user.Password != queryPassword {
		// c.IndentedJSON(http.StatusUnauthorized, 
		// HttpUnauthorized{Code: http.StatusUnauthorized, Status: "Unauthorized"})
		c.Status(http.StatusUnauthorized)
		return
	}
	// c.IndentedJSON(http.StatusOK, HttpOK{Code: http.StatusOK, Status: "OK"})
	c.Status(http.StatusOK)
}

func postUser(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	
	queryUsername := c.Query("username")
	queryPassword := c.Query("password")

	if queryUsername == "" || queryPassword == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	newUser := UserDocumentStruct{}
	newUser.Username = queryUsername
	newUser.Password = queryPassword
	newUser.Monitors = []string{""}

	_, errorInsert := insertUser(newUser)
	if errorInsert != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.IndentedJSON(http.StatusCreated, newUser)
}