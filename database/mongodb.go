package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/joho/godotenv"
)

var Collection *mongo.Collection

func MongoDB() *mongo.Database {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("can not load .env file", err)
	}

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB")))
	if err != nil {
		log.Fatal(err)
	}

	// status check
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	conn := client.Database(os.Getenv("MONGODB_DATABASE_NAME"))

	fmt.Println("MongoDB Connected")
	return conn
}
