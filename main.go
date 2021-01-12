/*---------------------------------------------------------------------------*\
 |  main.go                                                                  |
 |                                                                           |
 |  Copyright © 2020-2021, Rajiv Bakulesh Shah, original author.             |
 |  All rights reserved.                                                     |
\*---------------------------------------------------------------------------*/

package main

import (
	"log"
	"net/http"
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

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/v1/health", handleHealth)
	http.HandleFunc("/v1/gtfo", handleGTFO)
	http.HandleFunc("/v1/porn", handlePorn)

	port := ":" + getEnvVar("PORT")
	log.Println("Listening and serving HTTP on port " + port)
	err = http.ListenAndServe(port, nil)
	log.Fatal(err)
}
