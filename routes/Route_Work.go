package routes

import (
	"personal-web-be/handlers/work"

	"github.com/gofiber/fiber/v2"
)




func WorkRoute(app fiber.Router)  {
	workRoute := app.Group("/work-experience")


	workRoute.Get("/", work.GetWorkHandler)
	workRoute.Get("/:id", work.GetWorkHandlerById)
	workRoute.Post("/", work.CreateWorkHandler)
	workRoute.Delete("/:id", work.DeleteWorkHandler)
	workRoute.Patch("/:id", work.UpdateWorkHandler)
}