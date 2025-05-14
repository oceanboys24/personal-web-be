package middleware

import (
	model "personal-web-be/models"
	"personal-web-be/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetWIBTime() time.Time {
	wib := time.FixedZone("WIB", 7*3600)
	return time.Now().In(wib)
}



func ActiviyLogger() fiber.Handler  {
	uuidGen := uuid.New()
	return func (ctx *fiber.Ctx) error {
		method := ctx.Method()
        if method == "POST" || method == "PATCH" || method == "DELETE" {
		activity := model.Activity{
			ID: uuidGen.String(),
			Action: ctx.Method() ,
			Endpoint: ctx.OriginalURL(),
			CreatedAt: GetWIBTime(),
		}
		err := services.CreateLoggerActivity(activity)
		if err != nil {
			return ctx.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
				"message" : err.Error(),
			})
		}}
		return ctx.Next()
	}
}