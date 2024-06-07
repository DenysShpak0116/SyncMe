package controllers

import (
	"encoding/json"
	"net/http"
	"server/internal/database"
		"strconv"
	"github.com/go-chi/chi/v5"
)

func GetUserNotificationsFunc(w http.ResponseWriter, r *http.Request) {
	dbService := database.Instance()
	userId := chi.URLParam(r, "id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, "Cannot convert id to int: "+err.Error(), http.StatusBadRequest)
		return
	}

	notifications, err := dbService.GetUserNotifications(userIdInt)
	if err != nil {
		http.Error(w, "Cannot get user notifications: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(notifications)
}
