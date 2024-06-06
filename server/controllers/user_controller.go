package controllers

import (
	"encoding/json"
	"net/http"
	"server/internal/database"
)

func GetAllUsersFunc(w http.ResponseWriter, r *http.Request) {
	dbService := database.Instance()
	users, err := dbService.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func BanUserFunc(w http.ResponseWriter, r *http.Request) {
	dbService := database.Instance()
	var body struct {
		UserId int `json:"userId"`
	}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = dbService.ChangeUserRole(body.UserId, "banned")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func UnblockUserFunc(w http.ResponseWriter, r *http.Request) {
	dbService := database.Instance()
	var body struct {
		UserId int `json:"userId"`
	}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = dbService.ChangeUserRole(body.UserId, "user")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
