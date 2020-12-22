/*---------------------------------------------------------------------------*\
 |  porn.go                                                                  |
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
