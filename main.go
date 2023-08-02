package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Message       string    `json:"message"`
	Timestamp     time.Time `json:"timestamp"`
	UnixTimestamp int64     `json:"unix_timestamp"`
	SourceIP      string    `json:"src_ip"`
	Path          string    `json:"path"`
	Host          string    `json:"host"`
}

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	hostname, err := os.Hostname()

	if err != nil {
		hostname = "N/A"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		res, _ := json.Marshal(Response{
			Message:       "Hi",
			Timestamp:     time.Now(),
			UnixTimestamp: time.Now().Unix(),
			SourceIP:      r.RemoteAddr,
			Path:          r.RequestURI,
			Host:          hostname,
		})

		log.Printf("Request from %s to %s\n", r.RemoteAddr, r.RequestURI)

		w.Header().Add("Content-Type", "application/json")
		w.Write(res)

	})

	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
