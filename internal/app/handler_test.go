package app

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostShortUrl(t *testing.T) {

	type want struct {
		code        int
		response    string
		contentType string
	}

	tests := []struct {
		name string
		want want
	}{
		{
			name: "positive test POST",
			want: want{
				code:        201,
				response:    "http://localhost:8080/UwaF9CSP",
				contentType: "text/plain",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("https://practicum.yandex.ru/"))
			tRec := httptest.NewRecorder()
			PostShortURL(tRec, request)
			tRequestResult := tRec.Result()
			assert.Equal(t, tRequestResult.StatusCode, test.want.code)
			defer tRequestResult.Body.Close()
			_, err := io.ReadAll(tRec.Body)
			require.NoError(t, err)
			assert.Equal(t, tRequestResult.Header.Get("Content-Type"), test.want.contentType)
		})
	}
}

func TestPostShortUrl_Negative(t *testing.T) {

	type want struct {
		code int
		//response string
		// contentType string
	}

	tests := []struct {
		name string
		want want
	}{
		{
			name: "negative test POST",
			want: want{
				code: 400,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
			tRec := httptest.NewRecorder()
			if tRec.Body == nil {
				tRec.WriteHeader(http.StatusBadRequest)
			}
			PostShortURL(tRec, request)
			tRequestResult := tRec.Result()
			assert.Equal(t, tRequestResult.StatusCode, test.want.code)
			defer tRequestResult.Body.Close()
		})
	}
}

func TestGetOrigPageRedir(t *testing.T) {
	type want struct {
		code     int
		location string
	}

	tests := []struct {
		name string
		want want
	}{
		{
			name: "positive test GET",
			want: want{
				code:     307,
				location: "https://practicum.yandex.ru/",
			},
		},
	}
	for _, test := range tests {
		request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/UwaF9CSP", nil)
		resp := httptest.NewRecorder()
		GetOrigPageRedir(resp, request)
		respres := resp.Result()
		assert.Equal(t, test.want.code, respres.StatusCode)
		defer respres.Body.Close()
		assert.Equal(t, respres.Header.Get("Location"), test.want.location)
	}
}

func TestRootRouter_MethGet(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/UwaF9CSP", nil)
	resp := httptest.NewRecorder()
	RootRouter(resp, request)
	GetOrigPageRedir(resp, request)
}

func TestRootRouter_MethPost(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "/", nil)
	resp := httptest.NewRecorder()
	RootRouter(resp, request)

}

func TestRootRouter_BadReq(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "/.pkm", nil)
	resp := httptest.NewRecorder()
	RootRouter(resp, request)
}
