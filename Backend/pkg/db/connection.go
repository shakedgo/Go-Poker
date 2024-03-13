package db

import (
	"database/sql"
	"fmt"
	"log"
	// "github.com/lib/pq" // Import PostgreSQL driver
)

// Database holds the database connection.
type Database struct {
	*sql.DB
}

// NewDatabase initializes a new database connection.
func NewDatabase(host, port, user, password, dbname string) (*Database, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Check if the database connection is successful
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Connected to the database")

	return &Database{DB: db}, nil
}
