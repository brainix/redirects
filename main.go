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
				statusCode = http.StatusServiceUnavailable
				log.Println(err)
			}

			c.JSON(statusCode, gin.H{
				"message": http.StatusText(statusCode),
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
	router.Run(":" + port)
}
