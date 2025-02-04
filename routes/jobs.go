package routes

import (
	"encoding/json"
	"job-board/models"
	"math"
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
	// Extract job query parameter from the URL
	filterTitle := context.Query("query")

	// Extract pagination parameters with defaults
	page, err := strconv.Atoi(context.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1 // Default to page 1 if invalid
	}

	limit, err := strconv.Atoi(context.DefaultQuery("limit", "6"))
	if err != nil || limit < 1 {
		limit = 6 // Default to 6 items per page if invalid
	}

	// Get all jobs with filters and pagination
	jobs, total, err := models.GetAllJobs(filterTitle, page, limit)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch jobs: " + err.Error()})
		return
	}

	// Calculate total pages
	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	// Return jobs with the metadata(all jobs in the database & pagination)
	context.JSON(http.StatusOK, gin.H{
		"data": jobs,
		"metadata": gin.H{
			"current_page": page,
			"per_page":     limit,
			"total":        total,
			"total_pages":  totalPages,
		},
	})
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
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid job id."})
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
