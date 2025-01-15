package routes

import (
	"encoding/json"
	"job-board/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create job
func createJob(context *gin.Context) {
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
	// Extract query parameters
	filterTitle := context.Query("title")
	filterLocation := context.Query("location")

	// Call the model to fetch jobs with filters
	jobs, err := models.GetAllJobs(filterTitle, filterLocation)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch jobs: " + err.Error()})
		return
	}

	// Return the filtered jobs
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

// Delete a job
func deleteJob(context *gin.Context) {
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

	err = job.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete job"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "job deleted successfully!"})
}

// Update job
func updateJob(context *gin.Context) {
	// Parse job ID from the URL
	jobId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid job id"})
		return
	}

	// Parse the request body to get the updated job data
	var updatedJob models.Job
	if err := context.ShouldBindJSON(&updatedJob); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	// Serialize the Duties field to JSON for database storage
	dutiesJSON, err := json.Marshal(updatedJob.Duties)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error processing duties field"})
		return
	}

	// Update job in DB
	err = models.UpdateJobByID(jobId, updatedJob, string(dutiesJSON))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update job"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "job updated successfully"})
}
