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
