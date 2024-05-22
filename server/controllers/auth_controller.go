package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"server/internal/auth"
	"server/internal/database"
	"server/models"

	"github.com/go-chi/chi/v5"
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

	// Создание новой сессии
	session, _ := auth.Store.Get(r, auth.SessionName)
	session.Values["user"] = user
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

    res.Header().Set("Location", "/")
    res.WriteHeader(http.StatusTemporaryRedirect)
}

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    dbService := database.Instance() 
    err = dbService.AddUser(user)
    if err != nil {
        http.Error(w, "Could not register user", http.StatusInternalServerError)
        return
    }

	session, _ := auth.Store.Get(r, auth.SessionName)
	session.Values["username"] = user.Username
	session.Save(r, w)


    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}
