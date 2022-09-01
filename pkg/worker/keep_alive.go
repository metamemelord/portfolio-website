package worker

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var keepAliveBaseURL = "https://gaurav.dev"

func init() {
	keepAliveBaseURLfromEnv := os.Getenv("KEEP_ALIVE_BASE_URL")
	if len(keepAliveBaseURLfromEnv) > 0 {
		keepAliveBaseURL = keepAliveBaseURLfromEnv
	}
}

func KeepAlive(t time.Duration) {
	for {
		makeRequest()
		time.Sleep(t)
	}
}

func makeRequest() {
	res, err := http.Get(fmt.Sprintf("%s/health", keepAliveBaseURL))
	if err != nil {
		log.Println(err)
	} else {
		res.Body.Close()
	}
}
