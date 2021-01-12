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

func makeRedirectHandler(key string) func(w http.ResponseWriter, r *http.Request) {
	switch key {
	case "gtfo":
	case "porn":
	default:
		panic(`for makeRedirectHandler(), key must be "gtfo" or "porn"`)
	}
	redirectHandler := func(w http.ResponseWriter, r *http.Request) {
		url, err := client.SRandMember(key).Result()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.Unmarshal([]byte(url), &url)
		if key == "porn" {
			url = "https://www.reddit.com/" + url + "/"
		}
		http.Redirect(w, r, url, http.StatusFound)
	}
	return redirectHandler
}
