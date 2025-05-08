package routes

import (
	"personal-web-be/handlers/work"
	"personal-web-be/middleware"

	"github.com/gofiber/fiber/v2"
)




func WorkRoute(app fiber.Router)  {
	workRoute := app.Group("/work-experience")


	workRoute.Get("/", work.GetWorkHandler)
	workRoute.Get("/:id", work.GetWorkHandlerById)
	workRoute.Post("/",middleware.CheckToken, work.CreateWorkHandler)
	workRoute.Delete("/:id",middleware.CheckToken, work.DeleteWorkHandler)
	workRoute.Patch("/:id",middleware.CheckToken, work.UpdateWorkHandler)
}