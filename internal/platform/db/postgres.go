package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// New returns a ready-to-use *sql.DB connected to Postgres.
func New() (*sql.DB, error) {
	// Default DSN for local development
	defaultDSN := "postgres://roomunity:roomunity_password@localhost:5432/roomunity?sslmode=disable"

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		// Fallback to default DSN if env var is not set
		dsn = defaultDSN
	}

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("sql.Open error: %w", err)
	}

	// Basic connection pool tuning (adjust later if needed)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	// Check that the connection is actually working
	if err := db.Ping(); err != nil {
		// Check error to close and loger, no clean the last error
		if cerr := db.Close(); cerr != nil {
			log.Printf("error closing db after ping failure: %v", cerr)
		}
		return nil, fmt.Errorf("db.Ping error: %w", err)
	}

	return db, nil
}
