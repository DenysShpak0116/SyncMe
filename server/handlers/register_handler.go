package handlers

import (
	"encoding/json"
	"net/http"
	"server/models"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	
	var user models.User 

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, hashedPassword)
	if err != nil {
		http.Error(w, "Error saving user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}