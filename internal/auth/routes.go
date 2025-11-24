package auth

import (
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// RegisterRoutes registers all auth-related routes

type ServerDependencies interface {
	GetDatabaseClient() *mongo.Client
	GetRouter() *chi.Mux
}

func RegisterRoutes(server ServerDependencies) {
	router := server.GetRouter()
	dbClient := server.GetDatabaseClient()

	router.Post("/auth/signup", HandleSignup(dbClient))
	router.Post("/auth/login", HandleLogin)
}
