/*---------------------------------------------------------------------------*\
 |  health.go                                                                |
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
