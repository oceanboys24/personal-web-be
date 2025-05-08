package routes

import (
	"personal-web-be/handlers/stack"
	"personal-web-be/middleware"

	"github.com/gofiber/fiber/v2"
)


func StackRoute(app fiber.Router) {
	// Endpoint Stack
	stackRoute := app.Group("/stack")

	//Method Stack
	stackRoute.Get("/", stack.GetStackHandler)
	stackRoute.Post("/",middleware.CheckToken, stack.CreateStackHandler)
	stackRoute.Patch("/:id", middleware.CheckToken, stack.UpdateStackHandler)
	stackRoute.Delete("/:id",middleware.CheckToken, stack.DeleteStackHandler)
}	