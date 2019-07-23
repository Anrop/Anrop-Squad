package db

import (
	"database/sql"

	// Select SQL Driver
	_ "github.com/go-sql-driver/mysql"
)

var (
	// Database holds the connection to the database
	Database *sql.DB
)

// Connect to the specified database
func Connect(databaseURL string) error {
	var err error
	Database, err = sql.Open("mysql", databaseURL)
	return err
}
