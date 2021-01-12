/*---------------------------------------------------------------------------*\
 |  handlers.go                                                              |
 |                                                                           |
 |  Copyright © 2020-2021, Rajiv Bakulesh Shah, original author.             |
 |  All rights reserved.                                                     |
\*---------------------------------------------------------------------------*/

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handleHealth(w http.ResponseWriter, r *http.Request) {
	statusCode := http.StatusOK
	_, err := client.Ping().Result()
	if err != nil {
		log.Println(err)
		statusCode = http.StatusServiceUnavailable
	}
	message := http.StatusText(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(`{"message": "` + message + `"}`))
}

func handleGTFO(w http.ResponseWriter, r *http.Request) {
	url, err := client.SRandMember("gtfo").Result()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.Unmarshal([]byte(url), &url)
	http.Redirect(w, r, url, http.StatusFound)
}

func handlePorn(w http.ResponseWriter, r *http.Request) {
	subreddit, err := client.SRandMember("porn").Result()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.Unmarshal([]byte(subreddit), &subreddit)
	url := "https://www.reddit.com/" + subreddit + "/"
	http.Redirect(w, r, url, http.StatusFound)
}
