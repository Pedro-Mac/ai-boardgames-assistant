package main

import (
	"ai-assistant/boargames/db"
	"ai-assistant/boargames/routes"
	"ai-assistant/boargames/services"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

func main() {
	/* LOAD ENV VARIABLES */
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	/* DATABASE CONNECTION  */
	client, err := db.Connect()
	/* Initialize OpenAI client */
	openAiClient := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	embeddingService := services.NewEmbeddingService(openAiClient)

	if err != nil {
		panic(err)
	}

	router := chi.NewRouter()
	apiRouter := chi.NewRouter()

	router.Mount("/api/v1", apiRouter)

	/* ROUTING  */
	server := routes.NewServer(client, apiRouter, embeddingService)
	server.RegisterAllRoutes()

	http.ListenAndServe(":8080", router)
}
