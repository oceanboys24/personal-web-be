package activity

import (
	"personal-web-be/services"

	"github.com/gofiber/fiber/v2"
)


func GetActivityHandler(ctx *fiber.Ctx) error  {
	result, err  := services.GetLoggerActivity()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error" : "Cannot Get Activity" + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "Success Get Activity",
		"data" : result,
	})
}
