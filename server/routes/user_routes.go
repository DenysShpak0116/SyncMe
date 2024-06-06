package routes

import (
		"github.com/go-chi/chi/v5"
	"server/controllers"
)

func RegisterUserRoutes(r chi.Router) {
	r.Get("/allusers", controllers.GetAllUsersFunc)
	r.Post("/block", controllers.BanUserFunc)
	r.Post("/unblock", controllers.UnblockUserFunc)
}