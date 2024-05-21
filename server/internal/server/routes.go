package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/markbates/goth/gothic"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.HelloWorldHandler)

	r.Get("/health", s.healthHandler)

	r.Get("/auth/{provider}/callback", s.getAuthCallbackFuntion)
	r.Get("/auth/{provider}", s.getAuthFunction)
	r.Get("/logout/{provider}", s.logoutFunction)

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

func (s *Server) getAuthCallbackFuntion(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	r = r.WithContext(context.WithValue(context.Background(), "provider", provider));

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	fmt.Println(user)

	http.Redirect(w, r, "http://localhost:8080", http.StatusSeeOther)
}

func (s *Server) beginAuthProviderCallback(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider));
	gothic.BeginAuthHandler(w, r)
}

func (s *Server) getAuthFunction(res http.ResponseWriter, req *http.Request) {
	provider := chi.URLParam(req, "provider")
	req = req.WithContext(context.WithValue(context.Background(), "provider", provider))
	if gothUser, err := gothic.CompleteUserAuth(res, req); err == nil {
		fmt.Println(gothUser)
		http.Redirect(res, req, "http://localhost:8080", http.StatusSeeOther)
	} else {
		gothic.BeginAuthHandler(res, req)
	}
}

func(s *Server) logoutFunction(res http.ResponseWriter, req *http.Request) {
	gothic.Logout(res, req)
	res.Header().Set("Location", "/")
	res.WriteHeader(http.StatusTemporaryRedirect)
}