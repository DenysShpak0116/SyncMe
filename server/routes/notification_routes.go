package routes

import (
		"github.com/go-chi/chi/v5"
	"server/controllers"
)

func RegisterNotificationRoutes(r chi.Router) {
	r.Get("/notifications/{id}", controllers.GetUserNotificationsFunc)
}