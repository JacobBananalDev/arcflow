package config

import (
	"os"
)

// Config represents the application runtime configuration
// Values are sourced from enviornment variables to support
// containerized and production deployments
type Config struct {
	Port        string
	DatabaseURL string
}

// Load reads enviornment variables and constructs a Config instance.
// Defaults are applied only where appropriate, such as for the server port (e.g local dev port).
func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		// Default port for local development
		port = "8080"
	}

	dbURL := os.Getenv("DATABASE_URL")

	return &Config{
		Port: port,
		DatabaseURL: dbURL,
	}
}
