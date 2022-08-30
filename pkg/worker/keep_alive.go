package worker

import (
	"log"
	"net/http"
	"time"
)

var healthcheckURL = "https://gaurav.dev/health"

func KeepAlive(t time.Duration) {
	for {
		makeRequest()
		time.Sleep(t)
	}
}

func makeRequest() {
	res, err := http.Get(healthcheckURL)
	if err != nil {
		log.Println(err)
	} else {
		res.Body.Close()
	}
}
