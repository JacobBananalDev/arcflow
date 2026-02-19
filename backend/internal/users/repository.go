package users

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository provides database access for users.
type Repository struct {
	db *pgxpool.Pool
}

// NewRepository creates a new users repository.
func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

// Create inserts a new user into the database.
func (r *Repository) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (id, email, password_hash)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.Exec(ctx, query,
		user.ID,
		user.Email,
		user.PasswordHash,
	)

	return err
}

// GetByEmail retrieves a user by email.
func (r *Repository) GetByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT id, email, password_hash, created_at
		FROM users
		WHERE email = $1
	`

	row := r.db.QueryRow(ctx, query, email)

	var user User
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

