package main

import (
	"net/http"


	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/intermittent-reinforcement/shortener-proj/internal/app"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("route does not exist"))
	})

	r.Post("/", app.PostShortURL)
	r.Get("/{id}", app.GetOrigPageRedir)

	err := http.ListenAndServe(`:8080`, r)
	if err != nil {
		panic(err)
	}
}
