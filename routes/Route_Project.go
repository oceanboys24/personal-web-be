package routes

import (
	"personal-web-be/handlers/project"

	"github.com/gofiber/fiber/v2"
)





func ProjectRoute(app fiber.Router)  {
	routeProject := app.Group("/project")

	routeProject.Get("/", project.GetProjectHandler)
	routeProject.Post("/", project.CreateProjectHandler)
	routeProject.Delete("/:id", project.DeleteProjectHandler)
	routeProject.Patch("/:id", project.UpdateProjectHandler)
}