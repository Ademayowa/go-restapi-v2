package main

import (
	"os"

	"github.com/Ademayowa/go-restapi-v2/db"
	"github.com/Ademayowa/go-restapi-v2/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		println("No .env file found, using default environment variables")
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
