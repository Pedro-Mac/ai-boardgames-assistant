package routes

import (
	"ai-assistant/boargames/internal/auth"
	"ai-assistant/boargames/internal/boardgames"
	"ai-assistant/boargames/services"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Server struct {
	DatabaseClient   *mongo.Client
	Router           *chi.Mux
	EmbeddingService *services.EmbeddingService
}

func NewServer(dbClient *mongo.Client, router *chi.Mux, embeddingService *services.EmbeddingService) *Server {
	server := &Server{
		DatabaseClient:   dbClient,
		Router:           router,
		EmbeddingService: embeddingService,
	}

	return server
}

func (s *Server) GetDatabaseClient() *mongo.Client {
	return s.DatabaseClient
}

func (s *Server) GetRouter() *chi.Mux {
	return s.Router
}

func (s *Server) GetEmbeddingService() *services.EmbeddingService {
	return s.EmbeddingService
}

func (server *Server) RegisterAllRoutes() {
	auth.RegisterRoutes(server)
	boardgames.RegisterRoutes(server)
}
