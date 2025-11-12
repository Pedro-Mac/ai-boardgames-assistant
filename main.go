package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	/* ENTRY POINT */
	fmt.Println("Hello, World!")

	/*DATABASE CONNECTION */
	godotenv.Load()
	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		log.Fatal("The MongoDB URI is an empty string")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(clientOptions)

	if err != nil {
		panic(err)
	}

	// Sends a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	/* ROUTING  */

	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the home page!"))
	})

	router.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		var reqBody LoginRequestBody
		err := json.NewDecoder(r.Body).Decode(&reqBody)

		if err != nil {
			w.Write([]byte("Unable to login"))
			return
		}

	})

	http.ListenAndServe(":3000", router)
}

type LoginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
