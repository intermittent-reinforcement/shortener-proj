package main

import (
	//"fmt"
	"net/http"
	//"os"

	"github.com/go-chi/chi/v5"

	"github.com/intermittent-reinforcement/shortener-proj/internal/app/config"
	app "github.com/intermittent-reinforcement/shortener-proj/internal/app/handler"
)

func main() {

	config.NewConfig()

	r := chi.NewRouter()

	r.Post("/", app.PostShortURL)
	r.Post("/api/shorten", app.JSONShortURL)
	r.Get("/{id}", app.GetOrigPageRedir)

	baseURL := config.URLConfig.ServerAddress.Value
	err := http.ListenAndServe(baseURL, r)
	if err != nil {
		panic(err)
	}
}
