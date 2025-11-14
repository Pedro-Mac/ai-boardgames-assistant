package db

import (
	"log"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Connect() (*mongo.Client, error) {
	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		log.Fatal("The MongoDB URI is an empty string")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(clientOptions)

	if err != nil {
		log.Fatal("Connection to MongoDB failed:", err)
	}

	return client, nil
}
