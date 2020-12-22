/*---------------------------------------------------------------------------*\
 |  not_found.go                                                             |
 |                                                                           |
 |  Copyright © 2016-2020, Rajiv Bakulesh Shah, original author.             |
 |  All rights reserved.                                                     |
\*---------------------------------------------------------------------------*/

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func notFound(c *gin.Context) {
	statusCode := http.StatusNotFound
	name := http.StatusText(statusCode)
	c.JSON(http.StatusNotFound, gin.H{
		"status_code": statusCode,
		"name":        name,
	})
}
