package login

import (
	"encoding/json"
	"net/http"

	"github.com/matidev200/guardaloya-backend/internal/database"
	"github.com/matidev200/guardaloya-backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)


type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Response struct {
    Status  string `json:"status"`
    Message string `json:"message"`
    Token   string `json:"token,omitempty"` 
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginRequest
	var response Response

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	db := database.DB

	var dbUser models.User
	if err := db.Where("username = ?", loginRequest.Username).First(&dbUser).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(loginRequest.Password)); err != nil {
		http.Error(w, "Wrong credentials", http.StatusUnauthorized)
		return
	}

	tokenString, err := createToken(dbUser.Username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	response = Response{
		Status:  "success",
		Message: "Login Approved",
		Token:   tokenString,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}