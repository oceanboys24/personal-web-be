package routes

import (
	"personal-web-be/handlers/hero"

	"github.com/gofiber/fiber/v2"
)



func HeroRoute(app fiber.Router) {
	heroRoute := app.Group("/hero")


	heroRoute.Get("/", hero.HeroHandler)
}