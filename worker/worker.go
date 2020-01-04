package worker

import (
	"encoding/json"
	"log"
	"net/http"
)

type Data struct {
	GithubData    []interface{}
	WordpressData map[string]interface{}
}

var data = &Data{}

func GetData() *Data {
	return data
}

func RefreshData() {
	log.Println("Refreshing data")
	go githubPackageRefresher()
	go wordpressPostRefresher()
}

func githubPackageRefresher() {
	resp, err := http.Get("https://api.github.com/users/metamemelord/repos?per_page=100")
	if err != nil {
		log.Println("Error while reading from Github", err)
		return
	}
	githubResponse := []interface{}{}
	err = json.NewDecoder(resp.Body).Decode(&githubResponse)
	if err != nil {
		log.Println("Error while unmarshalling Github response", err)
		return
	}
	data.GithubData = githubResponse
}

func wordpressPostRefresher() {
	resp, err := http.Get("https://public-api.wordpress.com/rest/v1.1/sites/theanonymosopher.wordpress.com/posts/")
	if err != nil {
		log.Println("Error while reading from WordPress", err)
		return
	}
	wordpressResponse := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&wordpressResponse)
	if err != nil {
		log.Println("Error while unmarshalling WordPress response", err)
		return
	}
	data.WordpressData = wordpressResponse
}
