package register

import (
	"encoding/json"
	"net/http"

	"github.com/Easy-Job-Developer/catalog_plus/domain"
	"github.com/Easy-Job-Developer/catalog_plus/service"
)

func GetMyAllNames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allNames := service.GetAllNames()
	json.NewEncoder(w).Encode(allNames)
}

func CreateName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var name domain.User
	_ = json.NewDecoder(r.Body).Decode(&name)
	service.InsertOneName(name)
	json.NewEncoder(w).Encode(name)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	// Implement the functionality
}

func DeleteAName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	// Implement the functionality
}

func DeleteAllNames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := service.DeleteAllName()
	json.NewEncoder(w).Encode(count)
}

type SignUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var signUpReq SignUpRequest
	if err := json.NewDecoder(r.Body).Decode(&signUpReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate sign-up input
	signUpInput := domain.SignUpInput{Name: signUpReq.Name, Email: signUpReq.Email, Password: signUpReq.Password}
	if err := signUpInput.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create new user
	user := domain.User{Name: signUpReq.Name, Email: signUpReq.Email, Password: signUpReq.Password}
	if err := service.CreateUser(user); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "User created successfully"}
	json.NewEncoder(w).Encode(response)
}
