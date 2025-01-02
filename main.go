package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/jobs", getJobs)

	server.Run(":8080")
}

func getJobs(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
}
