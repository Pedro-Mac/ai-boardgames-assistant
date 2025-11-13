package main

import (
	"ai-assistant/boargames/db"
	"ai-assistant/boargames/routes"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	/* LOAD ENV VARIABLES */
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	/* DATABASE CONNECTION  */

	db.ConnectDB()

	if err != nil {
		panic(err)
	}

	/* ROUTING  */
	routes.Routes()
}
