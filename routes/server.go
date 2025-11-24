package routes

import (
	"ai-assistant/boargames/internal/auth"

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

func (s *Server) GetDatabaseClient() *mongo.Client {
	return s.DatabaseClient
}

func (s *Server) GetRouter() *chi.Mux {
	return s.Router
}

func (server *Server) RegisterAllRoutes() {
	auth.RegisterRoutes(server)
}
