package controller

import (
	"context"
	"fmt"
	"log"
	model "main/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

//mongodb helpers
func insertOneMovie(movie model.Netfix) {
	inserted,err := collection.InsertOne(context.Background(),movie)

	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("movie inserted with id", inserted.InsertedID)
}

//updateOne record
func updateOneMovie(movieId string){
	id,_:= primitive.ObjectIDFromHex(movieId)
	filter:=bson.M{"_id":id}
	update:=bson.M{"$set":bson.M{"watched":true}}

	result,err:= collection.UpdateOne(context.Background(),filter,update)

	if err!=nil{
		log.Fatal(err)
	}	
	fmt.Println("updated successfully",result.ModifiedCount)
}
