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

// Save user data into DB
func (job Job) Save() error {
	query := `
		INSERT INTO jobs(title, description, location, salary, job_id)
		VALUES(?, ?, ?, ?, ?)
	`

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
