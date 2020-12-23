/*---------------------------------------------------------------------------*\
 |  endpoints.go                                                             |
 |                                                                           |
 |  Copyright © 2016-2020, Rajiv Bakulesh Shah, original author.             |
 |  All rights reserved.                                                     |
\*---------------------------------------------------------------------------*/

package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func health(c *gin.Context) {
	statusCode := http.StatusOK
	_, err := client.Ping().Result()
	if err != nil {
		log.Println(err)
		statusCode = http.StatusServiceUnavailable
	}
	message := http.StatusText(statusCode)
	c.JSON(statusCode, gin.H{"message": message})
}

func gtfo(c *gin.Context) {
	url, err := client.SRandMember("gtfo").Result()
	if err == nil {
		url = url[1 : len(url)-1]
		c.Redirect(http.StatusFound, url)
	} else {
		log.Println(err)
		statusCode := http.StatusServiceUnavailable
		message := http.StatusText(statusCode)
		c.JSON(statusCode, gin.H{"message": message})
	}
}

func porn(c *gin.Context) {
	subreddit, err := client.SRandMember("porn").Result()
	if err == nil {
		subreddit = subreddit[1 : len(subreddit)-1]
		url := "https://www.reddit.com/" + subreddit + "/"
		c.Redirect(http.StatusFound, url)
	} else {
		log.Println(err)
		statusCode := http.StatusServiceUnavailable
		message := http.StatusText(statusCode)
		c.JSON(statusCode, gin.H{"message": message})
	}
}

func notFound(c *gin.Context) {
	statusCode := http.StatusNotFound
	name := http.StatusText(statusCode)
	c.JSON(http.StatusNotFound, gin.H{
		"status_code": statusCode,
		"name":        name,
	})
}
