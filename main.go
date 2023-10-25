package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imamsutono/go-gorm-postgresql/initializers"
)

var (
	server *gin.Engine
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variable", err)
	}

	initializers.ConnectDB(&config)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variable", err)
	}

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	log.Fatal(server.Run(":" + config.ServerPort))
}
