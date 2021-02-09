package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// https://blog.questionable.services/article/testing-http-handlers-go/
// https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server

func TestHome1(t *testing.T) {
	// Test request to home page
	res, err := http.Get("http://localhost:8080/")
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
}

func TestHome2(t *testing.T) {
	// Test request to a non-existent url
	// Create a request (returns a new incoming server request)
	request, _ := http.NewRequest("GET", "/abc", nil)

	// Record the response (returns an implementation of http.ResponseWriter
	// that records its mutations for later inspection
	response := httptest.NewRecorder()

	//
	home(response, request)

	got := response.Body.String()[:3]
	expected := "404"

	if got != expected {
		t.Errorf("Got %q, expected %q", got, expected)
	}

}

func TestHome3(t *testing.T) {
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
	if strings.Contains(received, expected) {
		t.Errorf("Body of response doesn't include '%s' as expected", expected)
	}
}
