package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/intermittent-reinforcement/shortener-proj/internal/app"
)

func main() {
	r := chi.NewRouter()
	r.Get("/{id}", app.GetOrigPageRedir)
	r.Post("/", app.PostShortURL)

	err := http.ListenAndServe(`:8080`, r)
	if err != nil {
		panic(err)
	}
}
