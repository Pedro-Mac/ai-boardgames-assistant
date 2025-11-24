package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func HandleSignup(dbClient *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody SignupRequestBody
		err := json.NewDecoder(r.Body).Decode(&reqBody)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Invalid request body",
			})
			return
		}

		authRepo := NewAuthRepository(dbClient, "ai-boardgame-assistant", "auth")
		result, err := authRepo.CreateUser(reqBody.Email, reqBody.Password)

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{
				"error": err.Error(),
			})
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"message": fmt.Sprintf("Inserted document with _id: %v\n", result.InsertedID),
		})
	}
}

func HandleLogin(dbClient *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody LoginRequestBody
		err := json.NewDecoder(r.Body).Decode(&reqBody)

		w.Header().Set("Content-Type", "application/json")

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Invalid request body",
			})
			return
		}

		authRepo := NewAuthRepository(dbClient, "ai-boardgame-assistant", "auth")

		authRepo.FindUserByCredentials(reqBody.Email, reqBody.Password)
	}
}
