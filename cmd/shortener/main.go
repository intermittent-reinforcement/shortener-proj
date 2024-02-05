package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/intermittent-reinforcement/shortener-proj/internal/app/config"
	"github.com/intermittent-reinforcement/shortener-proj/internal/app/handler"
)

func main() {

	config.NewConfig()

	r := chi.NewRouter()

	r.Post("/", handler.PostShortURL)
	r.Post("/api/shorten", handler.JSONShortURL)
	r.Get("/{id}", handler.GetOrigPageRedir)

	baseURL := config.URLConfig.ServerAddress.Value
	err := http.ListenAndServe(baseURL, r)
	if err != nil {
		log.Fatal(err)
	}
}
