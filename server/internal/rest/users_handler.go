package rest

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/dmytrodemianchuk/go-auth-mongo/internal/service"
)

type UsersHandler struct {
	usersService *service.Users
}

func NewUsersHandler(usersService *service.Users) *UsersHandler {
	return &UsersHandler{usersService: usersService}
}

func (h *UsersHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var inp service.SignUpInput
	if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.usersService.SignUp(r.Context(), inp); err != nil {
		http.Error(w, "SignUp failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UsersHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var inp service.SignInInput
	if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	token, err := h.usersService.SignIn(r.Context(), inp)
	if err != nil {
		http.Error(w, "SignIn failed", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"token": token}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *UsersHandler) Profile(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header is required", http.StatusUnauthorized)
		return
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenStr == "" {
		http.Error(w, "Bearer token is required", http.StatusUnauthorized)
		return
	}

	userID, err := h.usersService.ParseToken(r.Context(), tokenStr)
	if err != nil {
		http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
		return
	}

	user, err := h.usersService.GetUserByID(r.Context(), userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *UsersHandler) GetUserByName(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header is required", http.StatusUnauthorized)
		return
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenStr == "" {
		http.Error(w, "Bearer token is required", http.StatusUnauthorized)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Name parameter is required", http.StatusBadRequest)
		return
	}

	user, err := h.usersService.GetUserByName(r.Context(), name)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
