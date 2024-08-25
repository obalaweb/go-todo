package db

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connector() error {
	var err error
	// Attempt to open the database
	DB, err = sql.Open("sqlite3", "./db/app.db")
	if err != nil {
		return err
	}

	// Test the connection
	err = DB.Ping()
	if err != nil {
		return err
	}

	// Create the todos table if it doesn't exist
	_, err = DB.ExecContext(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS todos (
			id TEXT PRIMARY KEY,          
    		title TEXT NOT NULL,         
    		description TEXT,             
    		completed BOOLEAN NOT NULL,   
    		due_date TEXT,                
    		priority INTEGER,             
    		labels TEXT
		)`,
	)
	if err != nil {
		return err
	}

	// log.Println("Database connected and table created (if not exists).")
	return nil
}
