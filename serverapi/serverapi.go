package serverapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mattcullenmeyer/github-stargazer-server/repostars"
)

// https://tutorialedge.net/golang/creating-restful-api-with-golang/

// Define structure of API
type apiStruct struct {
	Repository string `json:"Repository"`
	Stars      string `json:"Stars"`
}

type apiError struct {
	Message string `json:"message"`
}

// API is a function to generate JSON string
// that tells how many stargazers a given GitHub repo has
func API(w http.ResponseWriter, r *http.Request) {

	// https://golangcode.com/get-a-url-parameter-from-a-request/

	// Get query parameters from url
	// (everything after "?")
	query := r.URL.Query()

	// Get query parameter(s) for "repo" key
	repos, ok := query["repo"]

	// Make sure query parameter(s) for "repo" key isn't missing
	if !ok || len(repos) == 0 {
		w.WriteHeader(http.StatusBadRequest) // return status code 400
		fmt.Fprintf(w, "Error: Url parameter 'repo' is missing.\n")
		fmt.Fprintf(w, "The url format should be http://localhost:8080/api?repo=<organization>/<repository>\n")
		fmt.Fprintf(w, "For example, the api call to get number of stargazers for https://github.com/mattcullenmeyer/anaplan should be:\n")
		fmt.Fprintf(w, "http://localhost:8080/api?repo=mattcullenmeyer/anaplan")
		return
	}

	// Query()["repo"] will return an array of items,
	// but we only want the first item
	repo := repos[0]

	// Run GetRepoStars function from repostars package,
	// which will make a request to the GitHub API
	// to get the number of stars for a given username/repository
	stars := repostars.GetRepoStars(repo)

	// Return error message if GitHub repository is not found
	if stars == "Not Found" {
		w.WriteHeader(http.StatusBadRequest) // return status code 400
		errorTxt := fmt.Sprintf("The following is not a valid public GitHub repository: https://github.com/%s", repo)
		//fmt.Fprintf(w, errorMessage)
		errorMessage := apiError{Message: errorTxt}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	// Format repository name and stargazer count as struct
	// to encode as JSON
	data := apiStruct{
		Repository: repo,
		Stars:      stars,
	}

	// Encode data array into a JSON string
	json.NewEncoder(w).Encode(data)
}
