package routes

import (
	"personal-web-be/handlers/visitor"

	"github.com/gofiber/fiber/v2"
)


func VisitorRoute(app fiber.Router) {
	// Route Visitor
	visitorRoute := app.Group("/visitor")

	// Route Handler
	visitorRoute.Get("/count", visitor.GetVisitorHandler)
	visitorRoute.Get("/", visitor.CountVisitorHandler)
}	