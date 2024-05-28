package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"server/internal/database"
	"server/models"
	"time"

	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt/v5"
)

func AddGroupFunc(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name                 string `json:"name"`
		GroupImage           string `json:"group_image"`
		GroupBackgroundImage string `json:"group_background_image"`
		Description 		 string `json:"description"`
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
		Description: 		body.Description,
	}
	groupId, err := dbService.AddGroup(group)
	if err != nil {
		http.Error(w, "Cannot add group: "+err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Group added successfully")

	cookie, err := r.Cookie("jwt-token")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	tokenString := cookie.Value
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	var user *models.User
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			http.Error(w, "Token is expired", http.StatusUnauthorized)
			return
		}

		user, err = dbService.GetUserById(int(claims["sub"].(float64)))
		if err != nil {
			http.Error(w, "Can not retrieve user "+err.Error(), http.StatusUnauthorized)
			return
		}
		log.Println("User retrieved successfully")
	}

	err = dbService.AddUserGroup(user.UserId, groupId)
	if err != nil {
		http.Error(w, "Cannot add user to group: "+err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("User added to group successfully")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]string{"message": "Group added successfully"}
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
