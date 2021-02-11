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
	Stars   int    `json:"stargazers_count"`
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

	// Parse body of response with github struct
	result := github{}
	jsonErr := json.Unmarshal(body, &result)
	// Log any errors
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	// Check to make sure url is valid
	// If "message": "Not Found" is returned
	// then return message (instead of number of stars)
	if result.Message == "Not Found" {
		return result.Message
	}

	// Convert Stars (type int) to type string
	stars := strconv.Itoa(result.Stars)
	// Return number of stars for given respository
	return stars
}
