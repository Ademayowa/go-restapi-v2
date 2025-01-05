package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/jobs", getJobs)
	server.POST("/jobs", createJob)
	server.GET("/jobs/:id", getJob)
}
