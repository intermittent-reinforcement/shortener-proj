package app

import (
	"io"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
)

const serverURL string = "http://localhost:8080"

var idMap = make(map[string]string)

func PostShortURL(res http.ResponseWriter, req *http.Request) {

	// Read original URL from request body
	bodyURL, err := io.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}
	origURL := string(bodyURL)
	// Verify Original URL validity
	_, err = url.ParseRequestURI(origURL)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	// Get ID for Original URL
	id := GenerateID(origURL)
	shortURL := serverURL + "/" + id
	// If ID does not exist - write that to idMap
	if _, exists := idMap[id]; !exists {
		idMap[id] = origURL
	}
	// Write HTTP responce status and headers
	res.Header().Set("Content-Type", "text/plain")

	res.WriteHeader(http.StatusCreated)
	// Write Shortened URL to HTTP response
	res.Write([]byte(shortURL))

}

func GetOrigPageRedir(res http.ResponseWriter, req *http.Request) {
	// Get hash from GET request URL Path
	hash := chi.URLParam(req, "id")
	// Check if hash exists in idMap (aka original URL is stored)
	origUrl, exists := idMap[hash]
	if exists {
		res.Header().Set("Location", origUrl)
		res.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		res.WriteHeader(http.StatusBadRequest)
	}
}
