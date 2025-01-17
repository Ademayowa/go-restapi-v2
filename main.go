package main

import (
	"job-board/db"
	"job-board/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		println("No .env file found. Using default environment variables")
	}

	db.InitDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":" + port)
}
