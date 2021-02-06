package repostars

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type github struct {
	Stars int `json:"stargazers_count"`
}

func GetRepoStars(url string) string {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	github1 := github{}
	jsonErr := json.Unmarshal(body, &github1)
	if jsonErr != nil {
		log.Fatal(err)
	}
	txt := strconv.Itoa(github1.Stars)
	return txt
}
