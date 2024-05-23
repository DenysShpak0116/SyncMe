package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"server/internal/database"
	"time"

	"github.com/golang-jwt/jwt"
)

func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				http.Error(w, "Token is expired", http.StatusUnauthorized)
				return
			}

			dbService := database.Instance()
			user, err := dbService.GetUserById(int(claims["sub"].(float64)))
			if err != nil {
				http.Error(w, "Can not retrieve user "+err.Error(), http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "user", user)
			r = r.WithContext(ctx)

			next(w, r)
		} else {
			fmt.Println(err)
		}

	}
}
