package models

import "job-board/db"

type Job struct {
	ID          int64
	Title       string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	Salary      string `binding:"required"`
	JobID       int64
}

// Save job into databse
func (job Job) Save() error {
	query := `
		INSERT INTO jobs(title, description, location, salary, job_id)
		VALUES(?, ?, ?, ?, ?)`

	sqlStmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer sqlStmt.Close()

	// Execute the SQL statement
	result, err := sqlStmt.Exec(job.Title, job.Description, job.Location, job.Salary, job.JobID)
	if err != nil {
		return err
	}

	// Add the auto generated ID from the database
	id, err := result.LastInsertId()
	job.ID = id

	return err
}

// Gell all jobs
func GetAllJobs() ([]Job, error) {
	query := "SELECT * FROM jobs"

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []Job
	for rows.Next() {
		var job Job

		// Read all columns from database
		err := rows.Scan(&job.ID, &job.Title, &job.Description, &job.Location, &job.Salary, &job.JobID)
		if err != nil {
			return nil, err
		}

		jobs = append(jobs, job)
	}

	return jobs, nil
}
