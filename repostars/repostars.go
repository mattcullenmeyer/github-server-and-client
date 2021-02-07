package repostars

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type github struct {
	Stars int `json:"stargazers_count"`
}

type notFound struct {
	Message string `json:"message"`
}

// GetRepoStars is a function to return number of stars for a given repo
func GetRepoStars(origin string) string {
	url := fmt.Sprintf("https://api.github.com/repos/%s", origin)

	// Get request to GitHub API
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	// Close body of request at end of function
	defer res.Body.Close()

	// Read the body of the request as bytes
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Parse body of response with nonFound struct
	found := notFound{}
	foundErr := json.Unmarshal(body, &found)
	// Log any errors
	if foundErr != nil {
		log.Fatal(foundErr)
	}
	// Check to make sure url is valid
	// If "message": "Not Found" is returned
	// then return "Not Found" (instead of number of stars)
	message := found.Message
	if message == "Not Found" {
		return message
	}

	// Parse body of response with github struct
	result := github{}
	jsonErr := json.Unmarshal(body, &result)
	// Log any errors
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	// Convert Stars (type int) to type string
	stars := strconv.Itoa(result.Stars)
	// Return number of stars for given respository
	return stars
}
