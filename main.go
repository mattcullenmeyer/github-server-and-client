package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mattcullenmeyer/github-server-and-client/serverapi"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "404 page not found.\n")
	fmt.Fprintf(w, "Url format should be http://localhost:8080/api?repo=<organization>/<repository>\n")
	fmt.Fprintf(w, "For example, the api call to get number of stargazers for https://github.com/mattcullenmeyer/anaplan should be:\n")
	fmt.Fprintf(w, "http://localhost:8080/api?repo=mattcullenmeyer/anaplan")
	return
}

func handleRequests() {
	// Create REST endpoint for API, mapping it to api function
	http.HandleFunc("/api", serverapi.API)
	// Handle all other paths and return meaningful 404 error
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
