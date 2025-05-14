package routes

import (
	"personal-web-be/handlers/activity"

	"github.com/gofiber/fiber/v2"
)


func ActivityRoute(app fiber.Router) {
	// Endpoint Activity
	activityRoute := app.Group("/activity")

	//Method Activity
	activityRoute.Get("/", activity.GetActivityHandler)
}	