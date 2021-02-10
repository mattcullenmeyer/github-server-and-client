package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// https://golang.org/pkg/net/http/httptest/
// https://blog.questionable.services/article/testing-http-handlers-go/
// https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server
// https://tutorialedge.net/golang/intro-testing-in-go/
// https://ieftimov.com/post/testing-in-go-testing-http-servers/

func TestHome1(t *testing.T) {
	// Test request to home page url
	res, err := http.Get("http://localhost:8080/")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("Received non-200 response: %d\n", res.StatusCode)
	}
	actual, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	expected := "404"
	received := string(actual)
	if !strings.Contains(received, expected) {
		t.Errorf("Body of response doesn't include '%s' as expected", expected)
	}
}

func TestHome2(t *testing.T) {
	// Test request to a non-existent url
	res, err := http.Get("http://localhost:8080/abc")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("Received non-200 response: %d\n", res.StatusCode)
	}
	actual, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	expected := "404"
	received := string(actual)
	if !strings.Contains(received, expected) {
		t.Errorf("Body of response doesn't include '%s' as expected", expected)
	}
}

// Use a table-driven approach to testing the handler
func TestHome(t *testing.T) {
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
			statusCode: http.StatusOK,
		},
		{
			name:       "home page",
			method:     http.MethodGet,
			url:        "/",
			want:       "404",
			statusCode: http.StatusOK,
		},
	}

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
				t.Errorf(received)
				//t.Errorf("Body of response doesn't include '%s' as expected", tc.want)
			}

		})
	}
}
