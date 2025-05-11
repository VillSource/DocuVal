package sqliteDB

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteService struct {
	db *sql.DB
}

// NewSQLiteService initializes a new SQLiteService
func NewSQLiteService(dbPath string) (*SQLiteService, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open SQLite database: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to SQLite database: %w", err)
	}

	return &SQLiteService{db: db}, nil
}

// Close closes the database connection
func (s *SQLiteService) Close() error {
	return s.db.Close()
}

// Execute executes a query without returning rows (e.g., INSERT, UPDATE, DELETE)
func (s *SQLiteService) Execute(query string, args ...interface{}) (sql.Result, error) {
	return s.db.Exec(query, args...)
}

// Query executes a query that returns rows
func (s *SQLiteService) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return s.db.Query(query, args...)
}

// QueryRow executes a query that returns a single row
func (s *SQLiteService) QueryRow(query string, args ...interface{}) *sql.Row {
	return s.db.QueryRow(query, args...)
}

// CreateTable creates a new table
func (s *SQLiteService) CreateTable(query string) error {
	_, err := s.Execute(query)
	return err
}

