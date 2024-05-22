package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"server/internal/auth"
	"server/internal/database"
	"server/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt/v5"
	"github.com/markbates/goth/gothic"
)

func GetAuthCallbackFuntion(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	session, _ := auth.Store.Get(r, auth.SessionName)
	session.Values["username"] = user.NickName
	session.Values["provider"] = provider
	session.Values["logged_in"] = true
	session.Save(r, w)

	fmt.Println(user)

	http.Redirect(w, r, "http://localhost:8080/#/", http.StatusSeeOther)
}

func GetAuthFunction(res http.ResponseWriter, req *http.Request) {
	provider := chi.URLParam(req, "provider")
	req = req.WithContext(context.WithValue(req.Context(), "provider", provider))
	if gothUser, err := gothic.CompleteUserAuth(res, req); err == nil {
		fmt.Println(gothUser)
		http.Redirect(res, req, "http://localhost:8080/#/", http.StatusSeeOther)
	} else {
		gothic.BeginAuthHandler(res, req)
	}
}

func LogoutFunction(res http.ResponseWriter, req *http.Request) {
	provider := chi.URLParam(req, "provider")
	if provider == "google" || provider == "facebook" {
		gothic.Logout(res, req)
	} else {
		fmt.Println("Logged out from non-api account")
	}

	session, _ := auth.Store.Get(req, auth.SessionName)
	session.Options.MaxAge = -1
	session.Save(req, res)

	http.Redirect(res, req, "http://localhost:8080", http.StatusSeeOther)

	res.Header().Set("Location", "/")
	res.WriteHeader(http.StatusTemporaryRedirect)
}

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Username  string `json:"username"`
		Password  string `json:"password"`
		Email     string `json:"email"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Sex       string `json:"sex"`
		Country   string `json:"country"`
	}

	if err := render.Decode(r, &body); err != nil {
		http.Error(w, "Cannot decode: "+err.Error(), http.StatusBadRequest)
		return
	}

	if body.Username == "" || body.Password == "" || body.Email == "" {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	hash, err := auth.HashPassword(body.Password)
	if err != nil {
		http.Error(w, "Cannot hash password: "+err.Error(), http.StatusInternalServerError)
		return
	}

	user := models.User{
		Username:  body.Username,
		Password:  hash,
		Email:     body.Email,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Sex:       body.Sex,
		Country:   body.Country,
		Role:      "user",
	}

	dbService := database.Instance()
	err = dbService.AddUser(user)
	if err != nil {
		http.Error(w, "Cannot add user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]string{"message": "User registered successfully"}
	json.NewEncoder(w).Encode(response)
}

func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := render.Decode(r, &body); err != nil {
		http.Error(w, "Cannot decode: "+err.Error(), http.StatusBadRequest)
		return
	}

	dbService := database.Instance()
	user, err := dbService.GetUserByUsername(body.Username)
	if err != nil {
		http.Error(w, "User not found: "+err.Error(), http.StatusNotFound)
		return
	}

	if !auth.CheckPasswordHash(body.Password, user.Password) {
		http.Error(w, "Password is not valid", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.UserId,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		http.Error(w, "Cannot sign token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt-token",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		HttpOnly: true,
		Path:     "/",
		Domain:   "localhost",
		Secure:   false, 
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"token": tokenString}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Cannot encode response: "+err.Error(), http.StatusInternalServerError)
	}
}

func Validate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	user := r.Context().Value("user").(*models.User)

	response := map[string]interface{}{
		"message": "Token is valid",
		"user":    user,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Cannot encode response: "+err.Error(), http.StatusInternalServerError)
	}
}