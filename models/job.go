package models

import (
	"encoding/json"
	"job-board/db"
	"strings"
)

// type Job struct {
// 	ID          int64
// 	Title       string   `binding:"required"`
// 	Description string   `binding:"required"`
// 	Location    string   `binding:"required"`
// 	Salary      string   `binding:"required"`
// 	Duties      []string `binding:"required"`
// }

type Job struct {
	ID          int64    `json:"id"`
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Location    string   `json:"location" binding:"required"`
	Salary      string   `json:"salary" binding:"required"`
	Duties      []string `json:"duties" binding:"required"`
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

// Get all jobs (optional filtering by title or location)
func GetAllJobs(filterTitle, filterLocation string) ([]Job, error) {
	// Base query
	query := "SELECT * FROM jobs WHERE 1=1"
	args := []interface{}{}

	// Add filtering by title
	if strings.TrimSpace(filterTitle) != "" {
		query += " AND LOWER(title) LIKE ?"
		args = append(args, "%"+strings.ToLower(filterTitle)+"%")
	}

	// Add filtering by location
	if strings.TrimSpace(filterLocation) != "" {
		query += " AND LOWER(location) LIKE ?"
		args = append(args, "%"+strings.ToLower(filterLocation)+"%")
	}

	// Execute the query
	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []Job

	// Parse rows into the Job struct
	for rows.Next() {
		var job Job
		var dutiesJSON string

		err := rows.Scan(&job.ID, &job.Title, &job.Description, &job.Location, &job.Salary, &dutiesJSON)
		if err != nil {
			return nil, err
		}

		// Convert the duties JSON string to a slice of strings
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

func (job Job) Delete() error {
	query := "DELETE FROM jobs WHERE id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(job.ID)
	return err
}
