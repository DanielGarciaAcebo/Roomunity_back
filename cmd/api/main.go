package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"roomunity_back/internal/platform/bootstrap"

	platformdb "roomunity_back/internal/platform/db"
	middleware "roomunity_back/internal/platform/http"
)

func main() {
	// Read port from environment variable, or use 8080 as default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Open database connection
	db, err := platformdb.New()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer func(db *sql.DB) {
		if err := db.Close(); err != nil {
			log.Printf("error closing database: %v", err)
		}
	}(db)

	log.Println("✅ Successfully connected to database")

	mux := bootstrap.BuildMux(db)

	handler := middleware.WithCORS(middleware.WithRequestLogging(mux))

	log.Printf("Server listening at http://localhost:%s\n", port)
	// Start HTTP server (only ONE ListenAndServe)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
