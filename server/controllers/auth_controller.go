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
	"server/internal/utils"
)

func GetAuthCallbackFuntion(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	dbService := database.Instance()

	users, err := dbService.GetAllUsers()
	if err != nil {
		http.Error(w, "Cannot get all users: "+err.Error(), http.StatusInternalServerError)
		return
	}
	var userExists bool = false
	for _, u := range users {
		if u.Email == user.Email {
			userExists = true
			break
		}
	}

	if userExists {
		userToLogin, err := dbService.GetUserByEmail(user.Email)
		if err != nil {
			http.Error(w, "Cannot get user by email: "+err.Error(), http.StatusInternalServerError)
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": userToLogin.UserId,
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
			HttpOnly: false,
			Path:     "/",
			Domain:   "syncme-client-f465c8129900.herokuapp",
			Secure:   false,
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := map[string]string{"token": tokenString}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Cannot encode response: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	var newUser *models.User
	bgLink, err := utils.GetRandomPhoto()
	if err != nil {
		http.Error(w, "Cannot get random photo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = dbService.AddUser(models.User{
		Username:  user.NickName,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Country:   user.Location,
		Sex: 	   "Other",
		Role:      "user",
		Logo:      user.AvatarURL,
		BgImage:   bgLink,
	})
	if err != nil {
		http.Error(w, "Cannot add user: "+err.Error(), http.StatusInternalServerError)
		return
	}
	newUser, err = dbService.GetUserByEmail(user.Email)
	if err != nil {
		http.Error(w, "Cannot get user by email: "+err.Error(), http.StatusInternalServerError)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": newUser.UserId,
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
		HttpOnly: false,
		Path:     "/",
		Domain:   "syncme-client-f465c8129900.herokuapp",
		Secure:   false,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"token": tokenString}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Cannot encode response: "+err.Error(), http.StatusInternalServerError)
	}
	if err != nil {
		http.Error(w, "Cannot add user: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(user)

	http.Redirect(w, r, "https://syncme-client-f465c8129900.herokuapp.com/#/", http.StatusSeeOther)
}

func GetAuthFunction(res http.ResponseWriter, req *http.Request) {
	provider := chi.URLParam(req, "provider")
	req = req.WithContext(context.WithValue(req.Context(), "provider", provider))
	if gothUser, err := gothic.CompleteUserAuth(res, req); err == nil {
		fmt.Println(gothUser)
		http.Redirect(res, req, "https://syncme-client-f465c8129900.herokuapp.com/#/", http.StatusSeeOther)
	} else {
		gothic.BeginAuthHandler(res, req)
	}
}

func LogoutFunction(res http.ResponseWriter, req *http.Request) {
	provider := chi.URLParam(req, "provider")
	fmt.Printf("Logging out provider: %s\n", provider)

	if provider == "google" || provider == "facebook" {
		err := gothic.Logout(res, req)
		if err != nil {
			fmt.Printf("Error logging out from provider: %s\n", err)
		} else {
			fmt.Println("Logged out from API provider")
		}
	} else {
		fmt.Println("Logged out from non-api account")
	}

	// Invalidate the JWT token cookie
	cookie := http.Cookie{
		Name:     "jwt-token",
		Value:    "",
		Expires:  time.Unix(0, 0), // Set to Unix epoch time
		MaxAge:   -1,              // MaxAge<0 means delete cookie now
		HttpOnly: false,
		Path:     "/",
		Domain:   "syncme-client-f465c8129900.herokuapp",
		Secure:   false,
	}
	http.SetCookie(res, &cookie)
	fmt.Println("JWT cookie invalidated")

	// Verify cookie is set in the response header
	receivedCookie, err := req.Cookie("jwt-token")
	if err != nil {
		fmt.Printf("Cookie not set: %s\n", err)
	} else {
		fmt.Printf("Cookie after invalidation: %v\n", receivedCookie)
	}

	// Redirect to home page after logout
	http.Redirect(res, req, "https://syncme-client-f465c8129900.herokuapp.com/#/", http.StatusSeeOther)
	fmt.Println("Redirected to home page")
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
	fmt.Println(user)

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
		HttpOnly: false,
		Path:     "/",
		Domain:   "syncme-client-f465c8129900.herokuapp",
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

    user, ok := r.Context().Value("user").(*models.User)
    if !ok || user == nil {
        http.Error(w, "User not found in context", http.StatusInternalServerError)
        return
    }

    response := map[string]interface{}{
        "message": "Token is valid",
        "user":    user,
    }

    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, "Cannot encode response: "+err.Error(), http.StatusInternalServerError)
    }
}
