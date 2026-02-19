package users

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// CreateUserRequest represents the registration payload.
type CreateUserRequest struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"strongpassword123"`
}

// CreateUserResponse represents the registration response.
type CreateUserResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// LoginRequest represents login payload.
type LoginRequest struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"strongpassword123"`
}

// LoginResponse represents successful login response.
type LoginResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}


// Handler provides HTTP handlers for users.
type Handler struct {
	repo *Repository
}

// NewHandler creates a new users handler.
func NewHandler(repo *Repository) *Handler {
	return &Handler{repo: repo}
}

// CreateUser handles POST /users requests.
// CreateUser godoc
// @Summary Register a new user
// @Description Creates a new user with email and password
// @Tags users
// @Accept json
// @Produce json
// @Param request body CreateUserRequest true "User registration payload"
// @Success 201 {object} CreateUserResponse
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /users [post]
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Basic validation
	if input.Email == "" || input.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	// Hash password server-side
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(input.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		http.Error(w, "Failed to process password", http.StatusInternalServerError)
		return
	}

	user := &User{
		ID:           uuid.New().String(),
		Email:        input.Email,
		PasswordHash: string(hashedPassword),
		CreatedAt:    time.Now(),
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	if err := h.repo.Create(ctx, user); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	// Never return password hash
	response := struct {
		ID        string    `json:"id"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
	}{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	json.NewEncoder(w).Encode(response)
}


// Login handles POST /login requests.
// Login godoc
// @Summary Authenticate user
// @Description Authenticates a user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login payload"
// @Success 200 {object} LoginResponse
// @Failure 400 {string} string "Invalid input"
// @Failure 401 {string} string "Invalid credentials"
// @Router /login [post]
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var input LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if input.Email == "" || input.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	user, err := h.repo.GetByEmail(ctx, input.Email)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Compare hashed password
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(input.Password),
	)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	response := struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	}{
		ID:    user.ID,
		Email: user.Email,
	}

	json.NewEncoder(w).Encode(response)
}

