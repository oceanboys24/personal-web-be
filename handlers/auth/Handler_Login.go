package auth

import (
	model "personal-web-be/models"

	"github.com/gofiber/fiber/v2"
)



func LoginHandler(ctx *fiber.Ctx) error {
	var reqLogin model.LoginModel

	
	if err := ctx.BodyParser(&reqLogin); err != nil{
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "invalid Request Body", 
		})
	}
	

	return ctx.JSON(fiber.Map{
		"message" : "Welcome To Login",
		"data" : reqLogin,	
	})
}