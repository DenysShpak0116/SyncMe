package server

import (
	"encoding/json"
	"log"
	"net/http"

	"server/internal/auth"
	"server/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	 _ "github.com/gorilla/sessions"
)
func (s *Server) RegisterRoutes() http.Handler {
    r := chi.NewRouter()
    r.Use(middleware.Logger)

    // CORS middleware for all routes
    r.Use(cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:8080"},
        AllowCredentials: true,
    }).Handler)

    r.Get("/", s.HelloWorldHandler)
    r.Get("/health", s.healthHandler)

    r.Post("/session", s.userSessionHandler)

    routes.RegisterAuthRoutes(r)

    return r
}



func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)

	http.Redirect(w, r, "http://localhost:8080", http.StatusSeeOther)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}
func (s *Server) userSessionHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := auth.Store.Get(r, auth.SessionName)

	if username, ok := session.Values["username"]; ok {
		// Пользователь вошел в систему
		json.NewEncoder(w).Encode(map[string]interface{}{
			"logged_in": true,
			"username":      username,
		})
	} else {
		// Пользователь не вошел в систему
		json.NewEncoder(w).Encode(map[string]interface{}{
			"logged_in": false,
		})
	}
}