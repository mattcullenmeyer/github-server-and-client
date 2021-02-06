package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type repoDataStruct struct {
	Repo  string `json:"Repository"`
	Stars string `json:"Stars"`
}

type repoDataMap []repoDataStruct

func api(w http.ResponseWriter, r *http.Request) {
	data := repoDataMap{
		repoDataStruct{
			Repo:  "Test Repo",
			Stars: "Test Stars",
		},
	}
	json.NewEncoder(w).Encode(data)
}

func handleRequests() {
	http.HandleFunc("/api", api)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
