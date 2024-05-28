package routes

import (
	"server/controllers"

	"github.com/go-chi/chi/v5"
)

func RegisterGroupRoutes(r chi.Router) {
	r.Route("/groups", func(r chi.Router) {
		r.Post("/add", controllers.AddGroupFunc)
		r.Get("/get", controllers.GetGroupsFunc)
	})
}
