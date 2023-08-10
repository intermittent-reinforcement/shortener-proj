package app

import (
	//	"fmt"
	"net/http"
	"net/http/httptest"

	//"strings"

	//"strings"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"
)

func TestPostShortUrl(t *testing.T) {

	handler := http.HandlerFunc(PostShortURL)

	srv := httptest.NewServer(handler)

	defer srv.Close()

	tests := []struct {
		name        string
		code        int
		method      string
		requestBody string
	}{
		{
			name:        "positive test POST",
			code:        201,
			method:      http.MethodPost,
			requestBody: "https://google.com/",
		},
		{
			name:        "negative PUT",
			code:        400,
			method:      http.MethodPut,
			requestBody: "",
		},
		{
			name:        "negative DELETE",
			code:        400,
			method:      http.MethodDelete,
			requestBody: "",
		},
	}
	for _, test := range tests {

		req := resty.New().R()
		req.Method = test.method
		req.URL = srv.URL
		//	fmt.Println(srv.URL)
		req.Body = test.requestBody
		resp, err := req.Send()

		assert.NoError(t, err, "error making HTTP request")
		assert.Equal(t, test.code, resp.StatusCode(), "Response code didn't match expected")

	}
}

func TestRootRouter(t *testing.T) {
	handler := http.HandlerFunc(RootRouter)

	srv := httptest.NewServer(handler)

	defer srv.Close()

	testCases := []struct {
		method string
	}{
		{method: http.MethodGet},
		{method: http.MethodPut},
		{method: http.MethodDelete},
		{method: http.MethodPost},
	}
	for _, tc := range testCases {
		t.Run(tc.method, func(t *testing.T) {

			req := resty.New().R()
			req.Method = tc.method
			req.URL = srv.URL
			//fmt.Println(req.URL)
			_, err := req.Send()
			if err != nil {
				panic(err)
			}
			assert.NoError(t, err, "error making HTTP request")
			assert.Equal(t, tc.method, req.Method, "error")
			//fmt.Println(idMap)
		})
	}
}

func TestGetOrigPageRedir(t *testing.T) {

	http.HandleFunc("/", PostShortURL)
	//fmt.Println(idMap)
	handler := http.HandlerFunc(GetOrigPageRedir)

	srv := httptest.NewServer(handler)

	defer srv.Close()

	tests := []struct {
		name string
		code int
		//location string
		method string
		reqURL string
	}{
		{
			name:   "positive test GET",
			code:   307,
			reqURL: "/TGktgzJr",
			//location: "https://google.com/",
			method: http.MethodGet,
		},
		{
			name: "negative PUT",
			code: 400,
			//location: "",
			reqURL: "/UwaF9CSP",
			method: http.MethodPut,
		},
		{
			name: "negative DELETE",
			code: 400,
			//location: "",
			reqURL: "/UwaF9CSP",
			method: http.MethodDelete,
		},
	}
	for _, test := range tests {
		req := resty.New().R()
		req.Method = test.method
		req.URL = srv.URL + test.reqURL
		//req.R().SetPathParams
		//fmt.Println(req.URL)
		resp, err := req.Send()
		assert.NoError(t, err, "error making HTTP request")
		//assert.Equal(t, test.location, resp.Header().Get("Location"), "ty loh")
		//fmt.Println(resp.Header().Get("Location"))
		assert.Equal(t, test.code, resp.StatusCode(), "Response code didn't match expected")
	}
}
