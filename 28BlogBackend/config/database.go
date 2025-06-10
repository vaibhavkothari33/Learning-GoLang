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

	mongoURI := os.Getenv("MONGO_URI")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))

	if err != nil {
		log.Fatal(err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal("Failed to  ping mongoDB", err)
	}
	fmt.Println("Connected to mongo DB Successfully")

	dbName := os.Getenv("DATABASE_NAME")
	DB = client.Database(dbName)

}

func GetCollection(collectionName string) *mongo.Collection {
	return DB.Collection(collectionName)
}
