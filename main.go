package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var hipache Hipache
var service *gin.Engine

func setupHipache() {
	conn, err := NewHipache("localhost:6379")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	hipache = conn
}

func setupMiddleware() {
	if os.Getenv("API_KEY") != "" {
		service.Use(RequireAuth)
	}
	service.Use(RequireHipache)
}

func setupEndpoints() {
	service.GET("/frontends", GetFrontends)
	service.POST("/frontends", CreateFrontend)
	service.GET("/frontends/:fe", GetBackends)
	service.POST("/frontends/:fe", CreateBackend)
	service.DELETE("/frontends/:fe", DeleteFrontend)
	service.POST("/flush", FlushFrontends)
}

func main() {
	service = gin.Default()

	setupHipache()
	setupMiddleware()
	setupEndpoints()

	defer hipache.Close()
	service.Run(":5000")
}
