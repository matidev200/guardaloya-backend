package credentials

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matidev200/guardaloya-backend/internal/database"
	"github.com/matidev200/guardaloya-backend/internal/models"
)


func CreateCredential(w http.ResponseWriter, r *http.Request)  {
    var credential models.Credential

    if err := json.NewDecoder(r.Body).Decode(&credential); err != nil {
        http.Error(w, "Error in decoding JSON: " + err.Error(), http.StatusBadRequest)
        return
    }

    if err := database.DB.Create(&credential).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(credential)
}


func UpdateCredential(w http.ResponseWriter, r *http.Request)  {
	crendentialId := mux.Vars(r)["id"]
	var credential models.Credential

	if err := database.DB.First(&credential, crendentialId).Error; err != nil {
		http.Error(w, "Error finding credential: " + err.Error(), http.StatusBadRequest)
	} 

	var updateData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Model(&credential).Updates(updateData).Error; err != nil {
		http.Error(w, "Error updating credential: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(credential)
}

func DeleteCredential(w http.ResponseWriter, r *http.Request)  {
	crendentialId := mux.Vars(r)["id"]
    var credential models.Credential

	if err := database.DB.First(&credential, crendentialId).Error; err != nil {
		http.Error(w, "Error finding credential: " + err.Error(), http.StatusBadRequest)

    }
    database.DB.Delete(&credential)

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(credential)
}

func GetCredentials(w http.ResponseWriter, r *http.Request)  {
	credentialFilter :=  r.URL.Query().Get("search")
	UserId := r.URL.Query().Get("user_id")
	var credentials []models.Credential

	if err := database.DB.Where("user_id = ? AND title LIKE ?", UserId, "%"+credentialFilter+"%").Find(&credentials).Error; err != nil {
		http.Error(w, "Error finding credentials: " + err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(credentials)
}

func GetCredential(w http.ResponseWriter, r *http.Request)  {
	credentialId := mux.Vars(r)["id"]
	UserId := r.URL.Query().Get("user_id")
	var credentials models.Credential

	if err := database.DB.Where("id = ? AND user_id = ?", credentialId, UserId).First(&credentials).Error; err != nil {
		http.Error(w, "Error finding credentials: " + err.Error(), http.StatusBadRequest)
	}
	

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(credentials)
}