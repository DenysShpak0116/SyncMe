package routes

import (
	"server/controllers"
	"server/middleware"

	"github.com/go-chi/chi/v5"
)

func RegisterAuthRoutes(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Get("/{provider}/callback", controllers.GetAuthCallbackFuntion)
		r.Get("/{provider}", controllers.GetAuthFunction)
		r.Post("/register", controllers.RegisterUserHandler)
		r.Post("/login", controllers.LoginUserHandler)
		r.Get("/logout/{provider}", controllers.LogoutFunction)
	})

	r.Get("/validate", middleware.RequireAuth(controllers.Validate))
}
