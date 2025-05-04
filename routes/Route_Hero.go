package routes

import (
	"personal-web-be/handlers/hero"
	"personal-web-be/middleware"

	"github.com/gofiber/fiber/v2"
)



func HeroRoute(app fiber.Router) {
	heroRoute := app.Group("/hero")


	heroRoute.Patch("/",middleware.CheckToken, hero.HeroHandlerUpdate)
	heroRoute.Get("/", hero.HeroHandlerGet)
	heroRoute.Get("/landing", hero.HeroHandlerGet)
}