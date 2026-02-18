package db

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Connect initializes a PostgreSQL connection pool
// It validates connectivity before returning to ensure the
// application fails fast if the database is unreachable
func Connect(databaseURL string) (*pgxpool.Pool, error) {
	// Create a context with timeout to avoid hanging startup
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return nil, err
	}

	// Verify the database connectivity 
	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	log.Println("Database connected successfully")

	return pool, nil
}
