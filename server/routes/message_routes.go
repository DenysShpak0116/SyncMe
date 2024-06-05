package routes

import (
	"server/controllers"

	"github.com/go-chi/chi/v5"
)

func RegisterMessageRoutes(r chi.Router) {
	r.Route("/messages", func(r chi.Router) {
		r.Post("/add", controllers.AddMessageFunc)
		r.Post("/get", controllers.GetMessageFunc)
		r.Post("/delete", controllers.DeleteMessageFunc)
	})
}
