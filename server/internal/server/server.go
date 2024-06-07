// server/server.go

package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"server/internal/database"

	_ "github.com/gorilla/sessions"
	"github.com/rs/cors"
)

// Server struct to hold server configurations
type Server struct {
	port int
	db   database.Service
}

// NewServer initializes and returns a new server instance
func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,
		db:   database.New(),
	}

	// Создаем маршруты
	handler := NewServer.RegisterRoutes()

	// Создаем обработчик CORS с нужными заголовками
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://syncme-client-f465c8129900.herokuapp.com"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	}).Handler(handler)

	// Создаем сервер с обработчиком CORS
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      corsHandler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
