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

// GetRepoStars is a function to return number of stars for a given repo
func GetRepoStars(origin string) string {
	url := fmt.Sprintf("https://api.github.com/repos/%s", origin)
	fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	result := github{}
	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		log.Fatal(err)
	}
	txt := strconv.Itoa(result.Stars)
	return txt
}
