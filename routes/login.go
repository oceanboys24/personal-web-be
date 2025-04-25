package routes

import (
	"personal-web-be/handlers/auth"

	"github.com/gofiber/fiber/v2"
)


func LoginRoute(app fiber.Router) {
	loginRoute := app.Group("/login")


	loginRoute.Post("/", auth.LoginHandler)
}