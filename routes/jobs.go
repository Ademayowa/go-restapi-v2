package routes

import (
	"job-board/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create job
func createJob(context *gin.Context) {
	// Extract job data
	var job models.Job

	err := context.ShouldBindJSON(&job)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data" + err.Error()})
		return
	}

	// Call Save() for saving job into database
	job.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "job created", "job": job})
}

// Fetch all jobs
func getJobs(context *gin.Context) {
	jobs, err := models.GetAllJobs()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch jobs" + err.Error()})
		return
	}

	context.JSON(http.StatusOK, jobs)
}
