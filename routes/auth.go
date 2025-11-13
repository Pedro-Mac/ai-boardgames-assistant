package routes

import (
	"ai-assistant/boargames/internal/auth"

	"github.com/go-chi/chi/v5"
)

func AuthRoutes(router *chi.Mux) {
	router.Post("/login", auth.HandleLogin)
}
