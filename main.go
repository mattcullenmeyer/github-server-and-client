package main

import (
	"encoding/json"
	"fmt"
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
	data := repoDataMap{
		repoDataStruct{
			Repo:  "Test Repo",
			Stars: "Test Stars",
		},
	}
	// encode data array into a JSON string
	json.NewEncoder(w).Encode(data)
}

func testPage(w http.ResponseWriter, r *http.Request) {
	url := "https://api.github.com/repos/jasonrudolph/keyboard"
	txt := repostars.GetRepoStars(url)
	fmt.Fprintf(w, txt)
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	// Create REST endpoint for API, mapping it to api function
	http.HandleFunc("/test", testPage)
	http.HandleFunc("/api", api)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
