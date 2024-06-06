package routes

import (
	"server/controllers"

	"github.com/go-chi/chi/v5"
)

func RegisterMessageRoutes(r chi.Router) {
	r.Route("/messages", func(r chi.Router) {
		r.Post("/add", controllers.AddMessageFunc)
		r.Get("/chats/{id}", controllers.GetChatsFunc)
		r.Post("/get", controllers.GetMessageFunc)
		r.Delete("/delete", controllers.DeleteMessageFunc)
	})
}
