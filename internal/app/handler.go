package app

import (
	//"fmt"
	//"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const serverURL string = "http://localhost:8080/"

var idMap = make(map[string]string)

func RootRouter(res http.ResponseWriter, req *http.Request) {

	// Check path and method to route to corresponding function
	if req.Method == http.MethodPost && req.URL.Path == "/" {
		PostShortURL(res, req)
	} else if req.Method == http.MethodGet && req.URL.Path == "/" {
		GetOrigPageRedir(res, req)
		//fmt.Println("LALALA2", res.Header())
	} else {
		res.WriteHeader(http.StatusBadRequest)
	}
}

func PostShortURL(res http.ResponseWriter, req *http.Request) {
	// // Read original URL from request body
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
	// // Get ID for Original URL
	id := GenerateID(origURL)
	serverURL := req.URL.String()
	shortURL := serverURL + id
	// // If ID does not exist - write that to idMap
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
	reqRawURLToString := req.URL.String()
	hash := strings.TrimPrefix(reqRawURLToString, "/")
	//fmt.Println(hash, " ", idMap)

	// Check if hash exists in idMap (aka original URL is stored)
	origURL, exists := idMap[hash]
	if exists {
		res.WriteHeader(http.StatusTemporaryRedirect)
		res.Header().Set("Location", origURL)
	} else {
		res.WriteHeader(http.StatusBadRequest)
	}
	//fmt.Println("LALALA1", res.Header())
}
