package handlers

import (
	"encoding/json"
	"net/http"
	"server/models"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var storedPassword string
	err = db.QueryRow("SELECT password FROM users WHERE username = ?", user.Username).Scan(&storedPassword)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}
