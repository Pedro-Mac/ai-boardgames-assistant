package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes() {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	apiRouter := chi.NewRouter()
	AuthRoutes(apiRouter)

	router.Mount("/api", apiRouter)

	http.ListenAndServe(":3000", router)
}
