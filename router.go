/*---------------------------------------------------------------------------*\
 |  router.go                                                                |
 |                                                                           |
 |  Copyright © 2016-2020, Rajiv Bakulesh Shah, original author.             |
 |  All rights reserved.                                                     |
\*---------------------------------------------------------------------------*/

package main

import "github.com/gin-gonic/gin"

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/v1")
	{
		api.GET("/health", health)
		api.GET("/gtfo", gtfo)
		api.GET("/porn", porn)
	}
	router.NoRoute(notFound)
	return router
}
