package users

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// Handler provides HTTP handlers for users.
type Handler struct {
	repo *Repository
}

// NewHandler creates a new users handler.
func NewHandler(repo *Repository) *Handler {
	return &Handler{repo: repo}
}

// CreateUser handles POST /users requests.
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email        string `json:"email"`
		PasswordHash string `json:"password_hash"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user := &User{
		ID:           uuid.New().String(),
		Email:        input.Email,
		PasswordHash: input.PasswordHash,
		CreatedAt:    time.Now(),
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	if err := h.repo.Create(ctx, user); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
