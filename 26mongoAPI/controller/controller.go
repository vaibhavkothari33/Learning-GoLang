package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vaibhavkothari33/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://vaibhavkothari50:iZqTxT76IgSuoh9k@gobackend.jbas9pp.mongodb.net/?retryWrites=true&w=majority"

const dbName = "netflix"
const colName = "watchlist"

// MI
var collection *mongo.Collection

// connect with mongo

func init() {
	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo DB connection Success")

	collection = client.Database(dbName).Collection(colName)

	// instance
	fmt.Println("collection instance is ready")
}

// mongo db helpers - file
// insert one record

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted one movie in db with id:", inserted.InsertedID)
}

func updateOneRecord(moviedId string) {
	id, err := primitive.ObjectIDFromHex(moviedId) // convert string into object id
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count : ", result.ModifiedCount)
	// .m and .d when ordered then .m
}

func deleteOneMovie(moviedId string) {
	id, err := primitive.ObjectIDFromHex(moviedId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	// option
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got delete with delete count: ", deleteCount)

}

// delete all records

func deleteAllMovies() int64 {
	filter := bson.D{}
	deleteResult, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("number of movies delete: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

func getOneMovie(moviedId string) (bson.M, error){
	id, err := primitive.ObjectIDFromHex(moviedId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	var movie bson.M
	err = collection.FindOne(context.Background(), filter).Decode(&movie)
	if err != nil {
		log.Fatal(err)
	}
	
	// fmt.Println("Single movie is : ", getcount)
	return movie,nil
}

func getAllMovies() []bson.M {
	filter := bson.D{}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	var movies []bson.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies

}

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)

}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var movie model.Netflix
	json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)

	json.NewEncoder(w).Encode(movie)

}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	updateOneRecord(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}
func GetOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	params := mux.Vars(r)
	movieId :=params["id"]

	movie,err := getOneMovie(movieId)
	if err != nil{
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(movie)
	// getOneMovie(params["id"])
	// json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovies()
	json.NewEncoder(w).Encode(count)

}
