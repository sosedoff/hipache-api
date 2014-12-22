package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

const VERSION = "0.1.0"

var hipache Hipache
var service *gin.Engine

func getEnv(name string, def string) string {
	val := os.Getenv(name)

	if val == "" {
		return def
	}

	return val
}

func setupHipache() {
	host := fmt.Sprintf(
		"%s:%s",
		getEnv("REDIS_HOST", "localhost"),
		getEnv("REDIS_PORT", "6379"),
	)

	conn, err := NewHipache(host)

	if err != nil {
		fmt.Println("Redis connection error:", err)
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

func main() {
	// Do not print extra debugging information
	gin.SetMode("release")

	service = gin.Default()

	setupHipache()
	setupMiddleware()
	setupEndpoints()

	defer hipache.Close()

	port := getEnv("PORT", "5000")
	fmt.Printf("Starting hipache-api on 0.0.0.0:%s\n", port)
	service.Run(":" + port)
}
