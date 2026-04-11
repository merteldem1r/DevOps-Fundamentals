package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/merteldem1r/DevOps-Fundamentals/3-Docker-Containerization/src/internal/handlers"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	globalHandler := handlers.NewGlobalHandler()

	r.Group(func(r chi.Router) {
		r.Route("/api/v1", func(r chi.Router) {
			r.Get("", globalHandler.Get)
			r.Get("/health", globalHandler.GetHealth)
		})
	})

	return r
}
