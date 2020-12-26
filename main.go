/*---------------------------------------------------------------------------*\
 |  main.go                                                                  |
 |                                                                           |
 |  Copyright © 2016-2020, Rajiv Bakulesh Shah, original author.             |
 |  All rights reserved.                                                     |
\*---------------------------------------------------------------------------*/

package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var client *redis.Client

func getEnvVar(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatal("$" + key + " must be set")
	}
	return value
}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.StaticFile("/loaderio-caa44a0090b3e2f28909941b0c7b7e9f.txt", "./static/loaderio-caa44a0090b3e2f28909941b0c7b7e9f.txt")
	api := router.Group("/v1")
	{
		routes := map[string]gin.HandlerFunc{
			"/health": handleHealth,
			"/gtfo":   handleGTFO,
			"/porn":   handlePorn,
		}
		for relativePath, handler := range routes {
			api.GET(relativePath, handler)
			api.HEAD(relativePath, handler)
		}
	}
	router.NoRoute(handleNotFound)
	return router
}

func main() {
	redisURL := getEnvVar("REDIS_URL")
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatal(err)
	}
	client = redis.NewClient(opts)
	router := initRouter()

	port := getEnvVar("PORT")
	router.Run(":" + port)
}
