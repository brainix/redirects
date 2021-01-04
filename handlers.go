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

	"github.com/gin-gonic/gin"
)

func handleHealth(c *gin.Context) {
	statusCode := http.StatusOK
	_, err := client.Ping().Result()
	if err != nil {
		log.Println(err)
		statusCode = http.StatusServiceUnavailable
	}
	message := http.StatusText(statusCode)
	c.JSON(statusCode, gin.H{"message": message})
}

func handleGTFO(c *gin.Context) {
	url, err := client.SRandMember("gtfo").Result()
	if err != nil {
		log.Println(err)
		statusCode := http.StatusServiceUnavailable
		message := http.StatusText(statusCode)
		c.JSON(statusCode, gin.H{"message": message})
		return
	}

	json.Unmarshal([]byte(url), &url)
	c.Redirect(http.StatusFound, url)
}

func handlePorn(c *gin.Context) {
	subreddit, err := client.SRandMember("porn").Result()
	if err != nil {
		log.Println(err)
		statusCode := http.StatusServiceUnavailable
		message := http.StatusText(statusCode)
		c.JSON(statusCode, gin.H{"message": message})
		return
	}

	json.Unmarshal([]byte(subreddit), &subreddit)
	url := "https://www.reddit.com/" + subreddit + "/"
	c.Redirect(http.StatusFound, url)
}

func handleNotFound(c *gin.Context) {
	statusCode := http.StatusNotFound
	name := http.StatusText(statusCode)
	c.JSON(http.StatusNotFound, gin.H{
		"status_code": statusCode,
		"name":        name,
	})
}
