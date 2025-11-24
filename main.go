package main

import (
	"ai-assistant/boargames/db"
	"ai-assistant/boargames/routes"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	/* LOAD ENV VARIABLES */
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	/* DATABASE CONNECTION  */
	client, err := db.Connect()

	if err != nil {
		panic(err)
	}

	router := chi.NewRouter()
	apiRouter := chi.NewRouter()

	router.Mount("/api/v1", apiRouter)

	/* ROUTING  */
	server := routes.NewServer(client, apiRouter)
	server.RegisterAllRoutes()

	http.ListenAndServe(":8080", router)
}
