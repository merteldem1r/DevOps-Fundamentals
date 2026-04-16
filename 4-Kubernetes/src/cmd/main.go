package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/ping"))

	r.Get("/api/v1/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Health!"))
	})
	http.ListenAndServe(":8080", r)
}
