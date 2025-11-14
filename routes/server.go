package routes

import (
	"ai-assistant/boargames/internal/auth"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Server struct {
	DatabaseClient *mongo.Client
	Router         *chi.Mux
}

func NewServer(dbClient *mongo.Client) *Server {
	router := chi.NewRouter()

	server := &Server{
		DatabaseClient: dbClient,
		Router:         router,
	}

	server.setupRouting()

	return server
}

const API_BASE_PATH = "/api"

func (server *Server) setupRouting() {
	apiRouter := chi.NewRouter()
	server.Router.Mount(API_BASE_PATH, apiRouter)

	apiRouter.Group(func(r chi.Router) {
		r.Get("/auth/me", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "Authenticated user info"}`))
		})
		r.Post("/auth/login", auth.HandleLogin)
	})

	http.ListenAndServe(":3000", server.Router)
}
