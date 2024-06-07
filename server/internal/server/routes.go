package server

import (
	"encoding/json"
	"log"
	"net/http"

	"server/internal/auth"
	"server/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/gorilla/sessions"
	"github.com/rs/cors"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(
		cors.New(cors.Options{
			AllowedOrigins:         []string{"http://localhost:8080"},
			AllowOriginRequestFunc: func(r *http.Request, origin string) bool { return true },
			AllowedMethods:         []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:         []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:         []string{"Link"},
			AllowCredentials:       true,
			OptionsPassthrough:     true,
			MaxAge:                 3599,
		}).Handler,
	)
	r.Use(withCORS)
	r.Get("/", s.HelloWorldHandler)
	r.Get("/health", s.healthHandler)
	r.Get("/session", s.userSessionHandler)

	routes.RegisterNotificationRoutes(r)
	routes.RegisterMessageRoutes(r)
	routes.RegisterAuthorRoutes(r)
	routes.RegisterGroupRoutes(r)
	routes.RegisterAuthRoutes(r)
	routes.RegisterUserRoutes(r)
	return r
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		next.ServeHTTP(w, r)
	})
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

	log.Println("session.Values[\"username\"]: ", session.Values["username"])
	if username, ok := session.Values["username"]; ok {
		log.Println(w, "User is logged in")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"logged_in": true,
			"username":  username,
			"provider":  session.Values["provider"],
		})
	} else {
		log.Println(w, "User is not logged in")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"logged_in": false,
		})
	}
}
