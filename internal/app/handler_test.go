package app

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestPostShortURL(t *testing.T) {
	type want struct {
		code      int
		shortLink string
	}

	tests := []struct {
		testName     string
		locationLink string
		want         want
	}{
		{
			"Positive test #1",
			"http://ebmb4oy4knent.net/bsotu8cwy2n",
			want{
				201,
				"http://localhost:8080/Q2XK4CGY",
			},
		},
		{
			"Positive test #2",
			"https://practicum.yandex.ru/",
			want{
				201,
				"http://localhost:8080/UwaF9CSP",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(tt.locationLink)))
			response := httptest.NewRecorder()

			r := chi.NewRouter()
			r.Post("/", PostShortURL)

			r.ServeHTTP(response, request)

			result := response.Result()
			assert.Equal(t, tt.want.code, result.StatusCode)

			defer result.Body.Close()
			shortenedLinkTestResult, err := io.ReadAll(result.Body)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, tt.want.shortLink, string(shortenedLinkTestResult))
		})
	}
}

func TestGetOrigPageRedir(t *testing.T) {

	type want struct {
		code     int
		location string
	}
	tests := []struct {
		testName string
		id       string
		want     want
	}{
		{
			"Positive test #1",
			"Q2XK4CGY",
			want{
				307,
				"http://ebmb4oy4knent.net/bsotu8cwy2n",
			},
		},
		{
			"Positive test #2",
			"UwaF9CSP",
			want{
				307,
				"https://practicum.yandex.ru/",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			idMap = make(map[string]string)
			idMap[tt.id] = tt.want.location
			request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/"+tt.id, nil)
			response := httptest.NewRecorder()

			r := chi.NewRouter()
			r.Get("/{id}", GetOrigPageRedir)

			r.ServeHTTP(response, request)

			result := response.Result()
			assert.Equal(t, tt.want.code, result.StatusCode)
			defer result.Body.Close()

			locationHeader := response.Header().Get("Location")
			assert.Equal(t, tt.want.location, locationHeader)
		})
	}
}
