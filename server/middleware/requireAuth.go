package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	_ "log"
	"net/http"
	"os"
	"server/internal/database"
	"time"

	"github.com/golang-jwt/jwt"
)

func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var body struct {
            Token string `json:"token"`
        }

        // Decode the JSON body to get the token
        if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
            http.Error(w, "Cannot decode token from request body: "+err.Error(), http.StatusBadRequest)
            return
        }

        tokenString := body.Token
        if tokenString == "" {
            http.Error(w, "Token is missing", http.StatusBadRequest)
            return
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return []byte(os.Getenv("JWT_SECRET")), nil
        })
        if err != nil {
            http.Error(w, "Cannot parse token: "+err.Error(), http.StatusUnauthorized)
            return
        }

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            if float64(time.Now().Unix()) > claims["exp"].(float64) {
                http.Error(w, "Token is expired", http.StatusUnauthorized)
                return
            }

            dbService := database.Instance()
            user, err := dbService.GetUserById(int(claims["sub"].(float64)))
            if err != nil {
                http.Error(w, "Cannot retrieve user: "+err.Error(), http.StatusUnauthorized)
                return
            }

            ctx := context.WithValue(r.Context(), "user", user)
            r = r.WithContext(ctx)
            next(w, r)
        } else {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
        }
    }
}
