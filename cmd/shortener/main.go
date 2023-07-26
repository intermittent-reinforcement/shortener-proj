package main

import (
	"net/http"

	"github.com/intermittent-reinforcement/shortener-proj/internal/app"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.RootRouter)
	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
