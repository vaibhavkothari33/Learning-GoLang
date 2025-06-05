package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/vaibhavkothari33/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://vaibhavkothari50:iZqTxT76IgSuoh9k@gobackend.jbas9pp.mongodb.net/"

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
//insert one record

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
