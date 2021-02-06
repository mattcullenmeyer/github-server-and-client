package main

import (
	"encoding/json"
	"fmt"
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

// Create map for API using above struct as type
//type repoDataMap []repoDataStruct

func api(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()

	repos, ok := query["repo"]

	if !ok || len(repos) == 0 {
		fmt.Println("repos not present")
	}

	data := []apiStruct{}
	for i := 0; i < len(repos); i++ {
		origin := repos[i]
		stars := repostars.GetRepoStars(origin)
		data = append(data, apiStruct{Repository: origin, Stars: stars})
	}

	// encode data array into a JSON string
	json.NewEncoder(w).Encode(data)
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request")

	query := r.URL.Query()

	repos, ok := query["repo"]

	if !ok || len(repos) == 0 {
		fmt.Println("repos not present")
	}

	for i := 0; i < len(repos); i++ {
		fmt.Println(repos[i])
	}
}

func handleRequests() {
	// Create REST endpoint for API, mapping it to api function
	http.HandleFunc("/api", api)
	http.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
