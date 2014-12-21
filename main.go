package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

const VERSION = "0.1.0"

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
	service.DELETE("/frontends/:fe/backend", DeleteBackend)
	service.POST("/flush", FlushFrontends)
}

func getServicePort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	return port
}

func main() {
	// Do not print extra debugging information
	gin.SetMode("release")

	service = gin.Default()

	setupHipache()
	setupMiddleware()
	setupEndpoints()

	defer hipache.Close()

	port := getServicePort()
	fmt.Printf("Starting hipache-api on 0.0.0.0:%s\n", port)
	service.Run(":" + port)
}
