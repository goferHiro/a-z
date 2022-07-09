package controllers

import (
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

const connectionString = "mongodb+srv://root:Nashaat@cluster0.uflhp.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

const dbName = "netflix"
const collectionName = "watchlist"

var collection *mongo.Collection

func init() {
	clientOption := option.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
  collection = client.Database(dbName).Collection(collectionName)
  fmt.Println("Connection instance running")
}
