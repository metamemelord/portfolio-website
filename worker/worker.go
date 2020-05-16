package worker

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/metamemelord/portfolio-website/model"
)

type Data struct {
	GithubData    []interface{}
	WordpressData []*model.Post
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
	defer resp.Body.Close()
	githubResponse := []interface{}{}
	err = json.NewDecoder(resp.Body).Decode(&githubResponse)
	if err != nil {
		log.Println("Error while unmarshalling Github response", err)
		return
	}
	data.GithubData = githubResponse
}

func wordpressPostRefresher() {
	resp, err := http.Get("https://public-api.wordpress.com/rest/v1.1/sites/metamemelord.wordpress.com/posts/")
	if err != nil {
		log.Println("Error while reading from WordPress", err)
		return
	}
	defer resp.Body.Close()
	wordpressResponse := &model.WordpressResponse{}
	err = json.NewDecoder(resp.Body).Decode(&wordpressResponse)
	customResponse := model.WordPressResponseToCustomResponse(*wordpressResponse)
	if err != nil {
		log.Println("Error while unmarshalling WordPress response", err)
		return
	}

	upb := len(customResponse)
	for idx, post := range customResponse {
		post.ID = upb - idx
	}
	data.WordpressData = customResponse
}
