package main

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
const DATABASE = "botmonitoring"
const URL = "mongodb://127.0.0.1:27017"


const COLLECTIONBOTACTIVITY = "botactivity"
const COLLECTIONUSER = "user"

func Client() (mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(URL))
	return *client, err
}