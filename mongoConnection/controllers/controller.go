package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://saiprasad22010642:Saiprasad%402412youtuvecluster.hykautx.mongodb.net/"
const dbName = "netflixgolang"
const colName = "watchlist"

var collection *mongo.Collection

// connect with mongodb
func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to db
	client,err:= mongo.Connect(context.TODO(), clientOption)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("mongo connected successfully")
	collection = client.Database(dbName).Collection(colName)

	fmt.Println("collection instance created")

}
