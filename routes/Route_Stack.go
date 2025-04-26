package routes

import (

	"personal-web-be/handlers/stack"

	"github.com/gofiber/fiber/v2"
)


func StackRoute(app fiber.Router) {
	// Endpoint Stack
	stackRoute := app.Group("/stack")

	//Method Stack
	stackRoute.Get("/", stack.GetStackHandler)
	stackRoute.Post("/", stack.CreateStackHandler)
	stackRoute.Delete("/:id", stack.DeleteStackHandler)
}	