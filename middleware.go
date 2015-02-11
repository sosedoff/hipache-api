package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

// Require valid api authentication token
func RequireAuth(c *gin.Context) {
	token := c.Request.FormValue("api_key")

	if token == "" {
		ErrorResponse(401, "Missing 'api_key' patameter", c)
		c.Abort()
		return
	}

	if token != os.Getenv("API_KEY") {
		ErrorResponse(401, "Invalid api key", c)
		c.Abort()
		return
	}

	c.Next()
}

// Require established redis connection in hipache client
func RequireHipache(c *gin.Context) {
	if hipache.Redis != nil {
		c.Next()
		return
	}

	ErrorResponse(400, "Connection is not set", c)
	c.Abort()
}
