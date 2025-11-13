package db

import (
	"log"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectDB() (*mongo.Client, error) {
	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		log.Fatal("The MongoDB URI is an empty string")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	return mongo.Connect(clientOptions)
}
