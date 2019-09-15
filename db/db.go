package db

import (
	"database/sql"
	"time"

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
	if err != nil {
		return err
	}

	Database.SetConnMaxLifetime(1 * time.Minute)
	Database.SetMaxOpenConns(4)
	Database.SetMaxIdleConns(2)

	return nil
}
