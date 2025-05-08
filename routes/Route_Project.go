package routes

import (
	"personal-web-be/handlers/project"
	"personal-web-be/middleware"

	"github.com/gofiber/fiber/v2"
)





func ProjectRoute(app fiber.Router)  {
	routeProject := app.Group("/project")

	routeProject.Get("/", project.GetProjectHandler)
	routeProject.Get("/:id", project.GetProjectHandler)
	routeProject.Post("/",middleware.CheckToken, project.CreateProjectHandler)
	routeProject.Delete("/:id",middleware.CheckToken, project.DeleteProjectHandler)
	routeProject.Patch("/:id",middleware.CheckToken, project.UpdateProjectHandler)
}