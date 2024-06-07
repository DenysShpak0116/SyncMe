package routes

import (
	"server/controllers"

	"github.com/go-chi/chi/v5"
)

func RegisterMailRoutes(r chi.Router) {
	r.Post("/sendMail", controllers.SendMailFunc)
	r.Post("/sendVerificationCode", controllers.SendVerificationCodeFunc)
}
