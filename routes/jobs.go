package routes

import (
	"job-board/models"
	"net/http"
	"strconv"

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

// Fetch a single job
func getJob(context *gin.Context) {
	jobId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse job id"})
		return
	}

	job, err := models.GetJobByID(jobId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch job"})
		return
	}

	context.JSON(http.StatusOK, job)
}
