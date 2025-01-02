package db

import (
	"database/sql"

	_ "github.com/glebarez/go-sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "job.db")
	if err != nil {
		panic("could not connect to database")

	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()
}

func createTable() {
	createJobsTable := `
	CREATE TABLE IF NOT EXISTS jobs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		job_id INTEGER,
		title TEXT NOT NULL,
		location TEXT NOT NULL,
		salary TEXT NOT NULL,
    description TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createJobsTable)
	if err != nil {
		panic("could not create jobs table" + err.Error())
	}
}
