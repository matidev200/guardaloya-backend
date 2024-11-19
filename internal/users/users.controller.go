package users

import (
	"encoding/json"
	"net/http"

	"github.com/matidev200/guardaloya-backend/internal/database"
	"github.com/matidev200/guardaloya-backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)


func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
    var user models.User

    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Error in decoding JSON: " + err.Error(), http.StatusBadRequest)
        return
    }
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

	if err != nil {
        http.Error(w, "Error hashing password: " + err.Error(), http.StatusInternalServerError)
        return
    }

	user.Password = string(hashedPassword)



    db := database.DB
	if err := db.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}
