/*---------------------------------------------------------------------------*\
 |  main.go                                                                  |
 |                                                                           |
 |  Copyright © 2016-2020, Rajiv Bakulesh Shah, original author.             |
 |  All rights reserved.                                                     |
\*---------------------------------------------------------------------------*/

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Setup route group for the API
	api := router.Group("/v1")
	{
		api.GET("/gtfo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status_code": http.StatusNotFound,
			"name":        http.StatusText(http.StatusNotFound),
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// Start and run the server
	router.Run(":" + port)
}
