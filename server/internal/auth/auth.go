package auth

import (
	"log"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/google"
)

const (
	key    = "SyncMeKey123"
	MaxAge = 0
	IsProd = false
	SessionName = "my-app-session" 
)

var Store *sessions.CookieStore

func NewAuth() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	googleSecret := os.Getenv("GOOGLE_SECRET")

	facebookClientID := os.Getenv("FACEBOOK_KEY")
	facebookSecret := os.Getenv("FACEBOOK_SECRET")

	Store = sessions.NewCookieStore([]byte(key))
	Store.MaxAge(MaxAge)

	Store.Options.Path = "/"
	Store.Options.HttpOnly = true
	Store.Options.Secure = IsProd

	gothic.Store = Store

	goth.UseProviders(
		google.New(googleClientID, googleSecret, "http://localhost:3000/auth/google/callback"),
		facebook.New(facebookClientID, facebookSecret, "http://localhost:3000/auth/facebook/callback"),
	)
}
