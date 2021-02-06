package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mattcullenmeyer/github-server-and-client/repostars"
)

// https://tutorialedge.net/golang/creating-restful-api-with-golang/

// Create struct for API
type repoDataStruct struct {
	Repo  string `json:"Repository"`
	Stars string `json:"Stars"`
}

// Create map for API using above struct as type
type repoDataMap []repoDataStruct

func api(w http.ResponseWriter, r *http.Request) {
	origin := "mattcullenmeyer/github-server-and-client"
	stars := repostars.GetRepoStars(origin)
	data := repoDataMap{
		repoDataStruct{
			Repo:  origin,
			Stars: stars,
		},
	}
	// encode data array into a JSON string
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
