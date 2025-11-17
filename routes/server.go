package routes

import (
	"ai-assistant/boargames/internal/auth"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Server struct {
	DatabaseClient *mongo.Client
	Router         *chi.Mux
}

func NewServer(dbClient *mongo.Client, router *chi.Mux) *Server {

	server := &Server{
		DatabaseClient: dbClient,
		Router:         router,
	}

	return server
}

type SignupRequestBody struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (server *Server) HandleSignup() {
	server.Router.Post("/auth/signup", func(w http.ResponseWriter, r *http.Request) {
		var reqBody SignupRequestBody
		err := json.NewDecoder(r.Body).Decode(&reqBody)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Invalid request body",
			})
			return
		}

		authRepo := auth.NewAuthRepository(server.DatabaseClient, "ai-boardgame-assistant", "auth")
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
	})
}
