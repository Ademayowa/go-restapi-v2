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

// func getJobs(context *gin.Context) {
// 	// Search jobs by title or location
// 	filterTitle := context.Query("title")
// 	filterLocation := context.Query("location")

// 	jobs, err := models.GetAllJobs(filterTitle, filterLocation)

// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch jobs: " + err.Error()})
// 		return
// 	}

// 	context.JSON(http.StatusOK, jobs)
// }

// Testing production job filtering
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
