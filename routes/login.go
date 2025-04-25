package routes

import (
	"personal-web-be/handlers/auth"
	"github.com/gofiber/fiber/v2"
)


func LoginRoute(app fiber.Router) {
	// Endpoint Login
	loginRoute := app.Group("/login")

	//Method Login
	loginRoute.Post("/", auth.LoginHandler)
}	