package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
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

	// Create HTTP multiplexer (router)
	mux := http.NewServeMux()

	// Root handler (simple text to verify the server is running)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Roomunity_back listening on port " + port + "\n"))
		if err != nil {
			log.Printf("error writing response: %v", err)
		}
	})

	handler := middleware.WithCORS(mux)

	log.Printf("Server listening at http://localhost:%s\n", port)
	// Start HTTP server (only ONE ListenAndServe)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
