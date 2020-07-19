package worker

import (
	"log"
	"net/http"
	"time"
)

func KeepAlive(t time.Duration) {
	for {
		makeRequest()
		time.Sleep(t)
	}
}

func makeRequest() {
	url := "https://gaurav-saini.herokuapp.com/health"
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	} else {
		res.Body.Close()
	}
}
