/*---------------------------------------------------------------------------*\
 |  main.go                                                                  |
 |                                                                           |
 |  Copyright © 2016-2020, Rajiv Bakulesh Shah, original author.             |
 |  All rights reserved.                                                     |
\*---------------------------------------------------------------------------*/

package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func main() {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		log.Fatal("$REDIS_URL must be set")
	}
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatal(err)
	}
	client := redis.NewClient(opts)
	ctx := context.Background()

	router := gin.Default()

	api := router.Group("/v1")
	{
		api.GET("/health", func(c *gin.Context) {
			statusCode := http.StatusOK
			_, err := client.Ping(ctx).Result()
			if err != nil {
				log.Println(err)
				statusCode = http.StatusServiceUnavailable
			}
			message := http.StatusText(statusCode)
			c.JSON(statusCode, gin.H{"message": message})
		})

		api.GET("/gtfo", func(c *gin.Context) {
			url, err := client.SRandMember(ctx, "gtfo").Result()
			if err == nil {
				url = url[1 : len(url)-1]
				c.Redirect(http.StatusFound, url)
			} else {
				statusCode := http.StatusServiceUnavailable
				message := http.StatusText(statusCode)
				c.JSON(statusCode, gin.H{"message": message})
			}
		})
	}

	router.NoRoute(func(c *gin.Context) {
		statusCode := http.StatusNotFound
		name := http.StatusText(statusCode)
		c.JSON(http.StatusNotFound, gin.H{
			"status_code": statusCode,
			"name":        name,
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router.Run(":" + port)
}
