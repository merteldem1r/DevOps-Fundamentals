package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/merteldem1r/DevOps-Fundamentals/3-Docker-Containerization/src/internal/config"
	"github.com/merteldem1r/DevOps-Fundamentals/3-Docker-Containerization/src/internal/handlers"
)

func NewRouter(cfg *config.Config) *chi.Mux {
	r := chi.NewRouter()

	globalHandler := handlers.NewGlobalHandler(cfg.MESSAGE)

	r.Group(func(r chi.Router) {
		// r.Use(middlewares.RequestLogger) own custon logger middleware
		r.Use(middleware.Logger) // chi logger middleware
		r.Use(middleware.Recoverer)

		r.Route("/api/v1", func(r chi.Router) {
			r.Get("/", globalHandler.Get)
			r.Get("/health", globalHandler.GetHealth)
		})
	})

	return r
}
