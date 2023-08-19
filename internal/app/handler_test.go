package app

import (
	//"fmt"
	"net/http"

	"io"
	"net/http/httptest"

	"strings"

	//"strings"
	"testing"
	//"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testRequest(t *testing.T, ts *httptest.Server, method, path string) (*http.Response, string) {
	req, err := http.NewRequest(method, ts.URL+path, nil)
	require.NoError(t, err)

	resp, err := ts.Client().Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	return resp, string(respBody)
}

func TestPostShortUrl(t *testing.T) {

	tests := []struct {
		name        string
		code        int
		method      string
		requestBody string
		contentType string
	}{
		{
			name:        "positive test POST",
			code:        201,
			method:      http.MethodPost,
			requestBody: "https://practicum.yandex.ru/",
			contentType: "text/plain",
		},
		{
			name:        "negative Post",
			code:        400,
			method:      http.MethodPost,
			requestBody: "hflwe",
			contentType: "",
		},
		// {
		// 	name:        "negative DELETE",
		// 	code:        400,
		// 	method:      http.MethodDelete,
		// 	requestBody: "",
		// 	contentType: "",
		// },
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(test.method, "/", strings.NewReader(test.requestBody))
			responseRecorder := httptest.NewRecorder()
			PostShortURL(responseRecorder, request)
			tRequestResult := responseRecorder.Result()
			assert.Equal(t, tRequestResult.StatusCode, test.code)
			defer tRequestResult.Body.Close()
			_, err := io.ReadAll(responseRecorder.Body)
			require.NoError(t, err)
			assert.Equal(t, tRequestResult.Header.Get("Content-Type"), test.contentType)
		})
	}
}

func TestRootRouter(t *testing.T) {
	testServer := httptest.NewServer(RootRouter())
	defer testServer.Close()

	testCases := []struct {
		method string
		url    string
	}{
		{method: http.MethodPost, url: "/"},
		{method: http.MethodGet, url: "/TGktgzJr"},
		{method: http.MethodPut, url: "/"},
		{method: http.MethodDelete, url: "/"},
	}
	for _, testCase := range testCases {
		resp, _ := testRequest(t, testServer, testCase.method, testCase.url)
		assert.Equal(t, testCase.method, resp.Request.Method)
	}
}

func TestGetOrigPageRedir(t *testing.T) {

	tests := []struct {
		name     string
		code     int
		location string
		method   string
		reqURL   string
	}{
		{
			name:     "positive test Get",
			code:     307,
			reqURL:   "http://localhost:8080/UwaF9CSP",
			location: "https://practicum.yandex.ru/",
			method:   http.MethodGet,
		},
		{
			name:     "negative Get",
			code:     400,
			reqURL:   "http://localhost:8080/fhifgln",
			location: "",
			method:   http.MethodGet,
		},
		{
			name:     "negative Get",
			code:     400,
			reqURL:   "http://localhost:8080/inluih",
			location: "",
			method:   http.MethodGet,
		},
	}
	for _, test := range tests {
		request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/UwaF9CSP", nil)
		resp := httptest.NewRecorder()
		GetOrigPageRedir(resp, request)
		respres := resp.Result()
		assert.Equal(t, test.code, respres.StatusCode)
		defer respres.Body.Close()
		assert.Equal(t, test.location, resp.Header().Get("Location"))
	}
}
