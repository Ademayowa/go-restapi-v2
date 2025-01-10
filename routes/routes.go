package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Testing production job filtering
func RegisterRoutes(server *gin.Engine) {
	// Apply CORS middleware
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://job-board-v3.vercel.app"}, // Allow frontend domain
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,           // Allow cookies or auth headers
		MaxAge:           12 * time.Hour, // Cache preflight for 12 hours
	}))

	// Define routes
	server.GET("/jobs", getJobs)
	server.POST("/jobs", createJob)
	server.GET("/jobs/:id", getJob)
}

// func RegisterRoutes(server *gin.Engine) {
// 	// Apply CORS middleware
// 	server.Use(cors.New(cors.Config{
// 		AllowOrigins:     []string{"https://job-board-v3.vercel.app"}, // Allow frontend domain
// 		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
// 		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
// 		ExposeHeaders:    []string{"Content-Length"},
// 		AllowCredentials: true,           // Set to true if you need cookies or authentication headers
// 		MaxAge:           12 * time.Hour, // Cache preflight request for 12 hours
// 	}))

// 	// Define routes
// 	server.GET("/jobs", getJobs)
// 	server.POST("/jobs", createJob)
// 	server.GET("/jobs/:id", getJob)
// }
