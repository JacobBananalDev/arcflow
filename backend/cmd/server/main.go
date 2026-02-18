package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jacobbananaldev/arcflow/internal/config"
	"github.com/jacobbananaldev/arcflow/internal/db"
)

// main is the application entry point.
// It loads configuration, initializes infrastructure (like database connections), and starts the HTTP server.
func main() {
	// Load runtime configuration from environment variables
	cfg := config.Load()

	// Ensure required configuration is provided
	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	// Initialize database connection pool
	pool, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer pool.Close()

	// Initialize HTTP router and define routes
	r := chi.NewRouter()

	// Health check endpoint to verify the server is running
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Arcflow API running"))
	})

	log.Printf("Server running on :%s\n", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
