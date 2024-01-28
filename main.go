// Package main is the root package managing the gin-gonic REST API.

// main.go
package main

import (
	"fmt"
	"os"

	"github.com/1liale/maze-backend/api/handlers"
	"github.com/1liale/maze-backend/api/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

var db *gorm.DB

func init() {
	// load environment variables from .env
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal(err)
	}
	db = middlewares.InitDB()
}

func main() {
	port := fmt.Sprintf(":%s", os.Getenv("API_PORT"))

	// init gin server
	router := gin.New()

	// middlewares
	router.Use(
		ginlogrus.Logger(logrus.New()),
		gin.Recovery(),
		middlewares.PropDBEnv(db),
		cors.Default(), // enable CORS for all origins with all HTTP requests allowed by default
	)

	// gets a system check on api health
	router.GET("/api-health", handlers.SystemCheck)

	router.Run(port)
}
