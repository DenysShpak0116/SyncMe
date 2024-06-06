package auth

import (
	"fmt"
	"log"
	"os"

	// "path/filepath"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/google"
)

const (
	key         = "SyncMeKey123"
	MaxAge      = 0
	IsProd      = false
	SessionName = "my-app-session"
)

var Store *sessions.CookieStore
func NewAuth() {
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file (auth): %v", err)
		}
	}
	

    googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
    googleSecret := os.Getenv("GOOGLE_SECRET")

    facebookClientID := os.Getenv("FACEBOOK_KEY")
    facebookSecret := os.Getenv("FACEBOOK_SECRET")

    // Get the app base URL from the environment
    baseURL := os.Getenv("BASE_URL")
    if baseURL == "" {
        log.Fatalf("BASE_URL not set")
    }

    Store = sessions.NewCookieStore([]byte(key))
    Store.MaxAge(MaxAge)

    Store.Options.Path = "/"
    Store.Options.HttpOnly = true
    Store.Options.Secure = IsProd

    gothic.Store = Store

    googleCallbackURL := fmt.Sprintf("%s/auth/google/callback", baseURL)
    facebookCallbackURL := fmt.Sprintf("%s/auth/facebook/callback", baseURL)

    goth.UseProviders(
        google.New(googleClientID, googleSecret, googleCallbackURL),
        facebook.New(facebookClientID, facebookSecret, facebookCallbackURL),
    )
}
