package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"context"

	_ "github.com/lib/pq"
)

// User represents the user model
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserRepository manages the database operations for User
type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) Create(ctx context.Context, user User) (User, error) {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	err := repo.db.QueryRowContext(ctx, query, user.Name, user.Email).Scan(&user.ID)
	return user, err
}

func (repo *UserRepository) Get(ctx context.Context, id int) (User, error) {
	var user User
	query := `SELECT id, name, email FROM users WHERE id=$1`
	err := repo.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email)
	return user, err
}

func (repo *UserRepository) Update(ctx context.Context, id int, user User) (User, error) {
	query := `UPDATE users SET name=$1, email=$2 WHERE id=$3 RETURNING id, name, email`
	err := repo.db.QueryRowContext(ctx, query, user.Name, user.Email, id).Scan(&user.ID, &user.Name, &user.Email)
	return user, err
}

func (repo *UserRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id=$1`
	_, err := repo.db.ExecContext(ctx, query, id)
	return err
}

// UserService handles business logic for Users
type UserService struct {
	repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, name, email string) (User, error) {
	return s.repo.Create(ctx, User{Name: name, Email: email})
}

func (s *UserService) GetUser(ctx context.Context, id int) (User, error) {
	return s.repo.Get(ctx, id)
}

func (s *UserService) UpdateUser(ctx context.Context, id int, name, email string) (User, error) {
	return s.repo.Update(ctx, id, User{Name: name, Email: email})
}

func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

// UserHandler defines HTTP endpoints for User API
type UserHandler struct {
	service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	createdUser, err := h.service.CreateUser(r.Context(), user.Name, user.Email)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	user, err := h.service.GetUser(r.Context(), id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	updatedUser, err := h.service.UpdateUser(r.Context(), id, user.Name, user.Email)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updatedUser)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := h.service.DeleteUser(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete user", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	// Database connection
	dbURL := "postgres://postgres:password@localhost:5432/userdb?sslmode=disable"
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// Initialize repository, service, and handler
	repo := NewUserRepository(db)
	service := NewUserService(repo)
	handler := NewUserHandler(service)

	// Set up HTTP routes
	http.HandleFunc("/create", handler.CreateUser)
	http.HandleFunc("/get", handler.GetUser)
	http.HandleFunc("/update", handler.UpdateUser)
	http.HandleFunc("/delete", handler.DeleteUser)

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
