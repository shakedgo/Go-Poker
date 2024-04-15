package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client
	Name   = "Go-Poker"
)

func Connect() {
	log.Println("Connecting DB...")
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	db := os.Getenv("DB")

	opts := options.Client().ApplyURI(db).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	Client = client

	// Send a ping to confirm a successful connection
	if err := client.Database(Name).RunCommand(context.TODO(), bson.D{{Key: "serverStatus", Value: 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Printf("\x1b[1;32mPinged your deployment.\nYou successfully connected to MongoDB!\x1b[0m\n")
}
