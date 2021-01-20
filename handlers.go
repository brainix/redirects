/*---------------------------------------------------------------------------*\
 |  handlers.go                                                              |
 |                                                                           |
 |  Copyright © 2020-2021, Rajiv Bakulesh Shah, original author.             |
 |  All rights reserved.                                                     |
\*---------------------------------------------------------------------------*/

package main

import (
	"encoding/json"
	"net/http"
	"path"
)

type HealthResponse struct {
	Message string `json:"message"`
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	statusCode := http.StatusOK
	_, err := client.Ping().Result()
	if err != nil {
		statusCode = http.StatusServiceUnavailable
	}
	message := http.StatusText(statusCode)
	responseStruct := HealthResponse{Message: message}
	responseBody, err := json.Marshal(responseStruct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(responseBody)
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	key := path.Base(r.URL.Path)
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
