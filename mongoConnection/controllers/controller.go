package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	model "main/models"
	"net/http"

	"github.com/gorilla/mux"
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

//deletion query
func deleteOneMovie(movieId string){
	id,_:= primitive.ObjectIDFromHex(movieId)
	filter:=bson.M{"_id":id}

	deleteCount,err:= collection.DeleteOne(context.Background(),filter)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("movie deleted",deleteCount.DeletedCount)
}
func deleteManyMovie(){
	deleteCount,err:= collection.DeleteMany(context.Background(),bson.D{{}},nil)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("movie deleted",deleteCount.DeletedCount)
}

//get all movies from db
func getAllMovies() []primitive.M{
	cursor,err:= collection.Find(context.Background(),bson.D{{}})
	if err!=nil{
		log.Fatal(err)
	}
	var movies[] primitive.M
	for cursor.Next(context.Background()){
		var movie bson.M
		err:=cursor.Decode(&movie)
		if err!=nil{
			log.Fatal(err)
		}
		movies=append(movies,movie)
	}

	defer cursor.Close(context.Background())
	return movies
}

//actuall cotrollers till now was mongodb helpers
func GetAllMoviesController(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencoded")

	allMovies:=getAllMovies()


	json.NewEncoder(w).Encode(allMovies)

}

func CreateMoviesController(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencoded")

	w.Header().Set("Allow-Control-Allow-Methods","POST")

	var movie model.Netfix
	json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)

	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatchedController(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencoded")

	w.Header().Set("Allow-Control-Allow-Methods","PUT")

	params:=mux.Vars(r)

	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}


