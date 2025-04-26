package routes

import (
	"personal-web-be/handlers/upload"
	"personal-web-be/middleware"

	"github.com/gofiber/fiber/v2"
)


func UploadRoute(app fiber.Router) {
	// Endpoint Login
	uploadRoute := app.Group("/upload")

	//Method Login
	uploadRoute.Post("/",middleware.CheckToken, upload.UploadHandler)
}	