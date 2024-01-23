package app

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/intermittent-reinforcement/shortener-proj/internal/app/config"

	"github.com/go-chi/chi/v5"
)

// Structs for marshalling and unmarshalling JSON
type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	Result string `json:"result"`
}

// Storage for hashes
var idMap = make(map[string]string)

func PostShortURL(res http.ResponseWriter, req *http.Request) {

	// Set HTTP header
	res.Header().Set("Content-Type", "text/plain")

	// Read original URL from request body
	bodyURL, err := io.ReadAll(req.Body)
	if err != nil {
		handleError(res)
		return
	}
	defer req.Body.Close()
	origURL := string(bodyURL)

	// Verify Original URL validity
	url, err := url.ParseRequestURI(origURL)
	if err != nil {
		handleError(res)
	}
	if !isUrl(url) {
		handleError(res)
		return
	}
	// Get ID for original URL
	id := GenerateID(origURL)

	// If ID does not exist - write that to idMap
	checkIdMapForExistingId(id, origURL)

	// Generate a short link for user
	shortURL := getShortenedLink(id)

	// Write HTTP responce status
	res.WriteHeader(http.StatusCreated)

	// Write Shortened URL to HTTP response
	res.Write([]byte(shortURL))
}

func JSONShortURL(res http.ResponseWriter, req *http.Request) {

	// Set HTTP header
	res.Header().Set("Content-Type", "application/json")

	// Create a struct for unmarshalling the request
	var request ShortenRequest

	// Read original URL from request body
	body, err := io.ReadAll(req.Body)
	if err != nil {
		handleError(res)
		return
	}
	defer req.Body.Close()

	// Unmarshall URL
	if err := json.Unmarshal(body, &request); err != nil {
		handleError(res)
		return
	}
	// Get ID for original URL
	id := GenerateID(request.URL)

	// If ID does not exist - write that to idMap
	checkIdMapForExistingId(id, request.URL)

	// Generate a short link for user
	shortURL := getShortenedLink(id)

	// Marshall short link to a struct
	marshalledResponse := ShortenResponse{Result: shortURL}

	bytes, err := json.Marshal(marshalledResponse)
	if err != nil {
		handleError(res)
	}
	// Write HTTP responce status
	res.WriteHeader(http.StatusCreated)

	// Write Shortened URL to HTTP response
	res.Write(bytes)
}

func GetOrigPageRedir(res http.ResponseWriter, req *http.Request) {

	// Get hash from GET request URL Path
	hash := chi.URLParam(req, "id")

	// Check if hash exists in idMap (aka original URL is stored)
	origURL, exists := idMap[hash]

	if exists {
		res.Header().Set("Location", origURL)
		res.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		//res.WriteHeader(http.StatusNotFound)
		handleError(res)
	}
}

// Checks URL validity
func isUrl(url *url.URL) bool {
	return url.Scheme != "" && url.Host != ""
}

// Checks if the ID is in storage
func checkIdMapForExistingId(id, origURL string) {
	if _, exists := idMap[id]; !exists {
		idMap[id] = origURL
	}
}

func handleError(res http.ResponseWriter) {
	res.WriteHeader(http.StatusBadRequest)
}

func getShortenedLink(id string) string {
	return config.URLConfig.BaseURL.Value + "/" + id
}
