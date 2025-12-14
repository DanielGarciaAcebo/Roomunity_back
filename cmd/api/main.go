package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	platformdb "roomunity_back/internal/platform/db"
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

	// Minimal HTTP handler for the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Roomunity_back listening on port %s\n", port)
		log.Println("written bytes:", n, "err:", err)
	})

	// Log where the server is running
	log.Printf("Server listening at http://localhost:%s\n", port)

	// Start HTTP server
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
