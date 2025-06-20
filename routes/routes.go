package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Apply CORS middleware
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://job-board-v3.vercel.app", "http://localhost:8080"}, // Allow frontend domain
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,           // Allow cookies or auth headers
		MaxAge:           12 * time.Hour, // Cache preflight for 12 hours
	}))

	// Define routes
	server.GET("/jobs", getJobs)
	server.POST("/jobs", createJob)

	// Specialized routes should come before parameterized routes
	server.GET("/jobs/recent", GetRecentJobs)
	server.GET("/jobs/highest-salary", GetHighestSalaryJobs)

	// Parameterized routes
	server.GET("/jobs/:id", getJob)
	server.DELETE("/jobs/:id", deleteJob)
	server.PUT("/jobs/:id", updateJob)

}
