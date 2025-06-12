package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
const mongoURI  = "mongodb+srv://vaibhavkothari50:h1KT4aTESTGOSj9y@flowwise.bbypa1t.mongodb.net/?retryWrites=true&w=majority&appName=flowWise"
	// mongoURI := os.Getenv("MONGO_URI")
	// if mongoURI == "" {
	// 	log.Fatal("MONGO_URI environment variable is not set")
	// }


	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}
	
	fmt.Println("Connected to MongoDB Successfully")

	dbName := os.Getenv("DATABASE_NAME")
	if dbName == "" {
		log.Fatal("DATABASE_NAME environment variable is not set")
	}
	
	DB = client.Database(dbName)
}

func GetCollection(collectionName string) *mongo.Collection {
	if DB == nil {
		log.Fatal("Database not initialized. Call ConnectDB() first.")
	}
	return DB.Collection(collectionName)
}