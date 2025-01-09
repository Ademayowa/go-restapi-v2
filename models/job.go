package models

import (
	"encoding/json"
	"job-board/db"
)

type Job struct {
	ID          int64
	Title       string   `binding:"required"`
	Description string   `binding:"required"`
	Location    string   `binding:"required"`
	Salary      string   `binding:"required"`
	Duties      []string `binding:"required"`
}

// Save job into databse
func (job Job) Save() error {
	query := `
		INSERT INTO jobs(title, description, location, salary, duties)
		VALUES(?, ?, ?, ?, ?)`

	sqlStmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer sqlStmt.Close()

	// Serialize the Duties field to JSON
	dutiesJSON, err := json.Marshal(job.Duties)
	if err != nil {
		return err
	}

	// Execute the SQL statement
	result, err := sqlStmt.Exec(job.Title, job.Description, job.Location, job.Salary, string(dutiesJSON))
	if err != nil {
		return err
	}

	// Add the auto generated ID from the database
	id, err := result.LastInsertId()
	job.ID = id

	return err
}

// Get all jobs with optional filtering by title or location
func GetAllJobs(filterTitle, filterLocation string) ([]Job, error) {
	query := "SELECT * FROM jobs WHERE 1=1"
	args := []interface{}{}

	if filterTitle != "" {
		query += " AND title LIKE ?"
		args = append(args, "%"+filterTitle+"%")
	}

	if filterLocation != "" {
		query += " AND location LIKE ?"
		args = append(args, "%"+filterLocation+"%")
	}

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []Job

	for rows.Next() {
		var job Job
		var dutiesJSON string

		err := rows.Scan(&job.ID, &job.Title, &job.Description, &job.Location, &job.Salary, &dutiesJSON)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal([]byte(dutiesJSON), &job.Duties)
		if err != nil {
			return nil, err
		}

		jobs = append(jobs, job)
	}

	return jobs, nil
}

// Get a single job
func GetJobByID(id int64) (*Job, error) {
	query := "SELECT * FROM jobs WHERE id =?"
	row := db.DB.QueryRow(query, id)

	var job Job
	var dutiesJSON string

	// Scan the result into variables
	err := row.Scan(&job.ID, &job.Title, &job.Description, &job.Location, &job.Salary, &dutiesJSON)
	if err != nil {
		return nil, err
	}

	// Deserialize Duties field from JSON to []string
	err = json.Unmarshal([]byte(dutiesJSON), &job.Duties)
	if err != nil {
		return nil, err
	}

	return &job, nil
}
