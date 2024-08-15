package rest

import (
	"encoding/json"
	"net/http"

	"github.com/dmytrodemianchuk/go-auth-mongo/internal/domain"
	"github.com/dmytrodemianchuk/go-auth-mongo/internal/service"
)

type UsersHandler struct {
	service *service.Users
}

func NewUsersHandler(service *service.Users) *UsersHandler {
	return &UsersHandler{service: service}
}

func (h *UsersHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var inp domain.SignUpInput
	if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := h.service.SignUp(r.Context(), inp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *UsersHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var inp domain.SignInInput
	if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	token, err := h.service.SignIn(r.Context(), inp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
