package routes

import (
	"server/controllers"
	"github.com/go-chi/chi/v5"
)

func RegisterAuthorRoutes(r chi.Router) {
	r.Route("/authors", func(r chi.Router) {
		r.Post("/add", controllers.AddAuthorFunc)
	})
}