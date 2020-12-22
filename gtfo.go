/*---------------------------------------------------------------------------*\
 |  gtfo.go                                                                  |
 |                                                                           |
 |  Copyright © 2016-2020, Rajiv Bakulesh Shah, original author.             |
 |  All rights reserved.                                                     |
\*---------------------------------------------------------------------------*/

package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func gtfo(c *gin.Context) {
	ctx := context.Background()
	url, err := client.SRandMember(ctx, "gtfo").Result()
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
