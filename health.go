/*---------------------------------------------------------------------------*\
 |  health.go                                                                |
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

func health(c *gin.Context) {
	statusCode := http.StatusOK
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Println(err)
		statusCode = http.StatusServiceUnavailable
	}
	message := http.StatusText(statusCode)
	c.JSON(statusCode, gin.H{"message": message})
}
