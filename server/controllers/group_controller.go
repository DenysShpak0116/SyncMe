package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"server/internal/database"
	"server/models"

	"github.com/go-chi/render"
)

func AddGroupFunc(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name                 string `json:"name"`
		GroupImage           string `json:"group_image"`
		GroupBackgroundImage string `json:"group_background_image"`
		Description          string `json:"description"`
		UserId               int    `json:"user_id"`
	}

	if err := render.Decode(r, &body); err != nil {
		http.Error(w, "Cannot decode: "+err.Error(), http.StatusBadRequest)
		return
	}

	dbService := database.Instance()
	group := models.Group{
		Name:                 body.Name,
		GroupImage:           body.GroupImage,
		GroupBackgroundImage: body.GroupBackgroundImage,
		Description:          body.Description,
	}
	groupId, err := dbService.AddGroup(group)
	if err != nil {
		http.Error(w, "Cannot add group: "+err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Group added successfully")

	var user *models.User
	user, err = dbService.GetUserById(body.UserId)
	if err != nil {
		http.Error(w, "Can not retrieve user "+err.Error(), http.StatusUnauthorized)
		return
	}
	log.Println("User retrieved successfully")

	err = dbService.AddUserGroup(user.UserId, groupId)
	if err != nil {
		http.Error(w, "Cannot add user to group: "+err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("User added to group successfully")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]int{"group_id": groupId}
	json.NewEncoder(w).Encode(response)
}

func GetGroupsFunc(w http.ResponseWriter, r *http.Request) {
	dbServer := database.Instance()
	groups := dbServer.GetAllGroups()
	response := groups
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Cannot encode response: "+err.Error(), http.StatusInternalServerError)
	}
}

func FollowGroupFunc(w http.ResponseWriter, r *http.Request) {
	var body struct {
		UserId  int `json:"user_id"`
		GroupId int `json:"group_id"`
	}

	if err := render.Decode(r, &body); err != nil {
		http.Error(w, "Cannot decode: "+err.Error(), http.StatusBadRequest)
		return
	}

	dbService := database.Instance()

	err := dbService.AddUserGroup(body.UserId, body.GroupId)
	if err != nil {
		http.Error(w, "Cannot add user to group: "+err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("User added to group successfully")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]string{"message": "User added to group successfully"}
	json.NewEncoder(w).Encode(response)
}
