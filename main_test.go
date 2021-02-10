package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mattcullenmeyer/github-stargazer-server/serverapi"
)

// https://golang.org/pkg/net/http/httptest/
// https://blog.questionable.services/article/testing-http-handlers-go/
// https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server
// https://tutorialedge.net/golang/intro-testing-in-go/
// https://ieftimov.com/post/testing-in-go-testing-http-servers/

func TestHome(t *testing.T) {
	// Use a table-driven approach to testing the handler
	tt := []struct {
		name       string
		method     string
		url        string
		want       string
		statusCode int
	}{
		{
			name:       "non-existent url",
			method:     http.MethodGet,
			url:        "/abc",
			want:       "404",
			statusCode: http.StatusNotFound,
		},
		{
			name:       "home page",
			method:     http.MethodGet,
			url:        "/",
			want:       "404",
			statusCode: http.StatusNotFound,
		},
	}

	// Loop through each subtest
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// Create a request (returns a new incoming server request)
			request := httptest.NewRequest(tc.method, tc.url, nil)

			// Record the response (returns an implementation of http.ResponseWriter
			// that records its mutations for later inspection)
			response := httptest.NewRecorder()

			// Run home function
			home(response, request)

			// Check to make sure the status code is as expected
			if response.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, response.Code)
			}

			// Convert the response body to a string
			received := response.Body.String()

			// Search for substring "want" in the body of the response
			if !strings.Contains(received, tc.want) {
				t.Errorf("Body of response doesn't include '%s' as expected", tc.want)
			}

		})
	}
}

func TestAPI(t *testing.T) {
	// Use a table-driven approach to testing the handler
	tt := []struct {
		name       string
		method     string
		url        string
		want       string
		statusCode int
	}{
		{
			name:       "missing url parameter",
			method:     http.MethodGet,
			url:        "/api",
			want:       "Error:",
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "invalid github repo",
			method:     http.MethodGet,
			url:        "/api?repo=abc/def",
			want:       "not a valid",
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "valid github repo",
			method:     http.MethodGet,
			url:        "/api?repo=mattcullenmeyer/anaplan",
			want:       "anaplan",
			statusCode: http.StatusOK,
		},
	}

	// Loop through each subtest
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// Create a request (returns a new incoming server request)
			request := httptest.NewRequest(tc.method, tc.url, nil)

			// Record the response (returns an implementation of http.ResponseWriter
			// that records its mutations for later inspection)
			response := httptest.NewRecorder()

			// Run home function
			serverapi.API(response, request)

			// Check to make sure the status code is as expected
			if response.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, response.Code)
			}

			// Convert the response body to a string
			received := response.Body.String()

			// Search for substring "want" in the body of the response
			if !strings.Contains(received, tc.want) {
				//t.Errorf("Body of response doesn't include '%s' as expected", tc.want)
				t.Errorf(received)
			}

		})
	}
}
