package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mattcullenmeyer/github-server-and-client/repostars"
)

// https://tutorialedge.net/golang/creating-restful-api-with-golang/

// Define structure of API
type apiStruct struct {
	Repository string `json:"Repository"`
	Stars      string `json:"Stars"`
}

func api(w http.ResponseWriter, r *http.Request) {

	// https://golangcode.com/get-a-url-parameter-from-a-request/
	query := r.URL.Query()

	repos, ok := query["repo"]

	if !ok || len(repos) == 0 {
		log.Println("Url Param 'repo' is missing")
		return
	}

	// Query()["repo"] will return an array of items,
	// but we only want the first item
	repo := repos[0]

	// Run GetRepoStars function from repostars package,
	// which will make a request to the GitHub API
	// to get the number of stars for a given username/repository
	stars := repostars.GetRepoStars(repo)
	data := apiStruct{
		Repository: repo,
		Stars:      stars,
	}

	// Encode data array into a JSON string
	json.NewEncoder(w).Encode(data)
}

func handleRequests() {
	// Create REST endpoint for API, mapping it to api function
	http.HandleFunc("/api", api)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
