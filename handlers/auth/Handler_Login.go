package auth

import (
	model "personal-web-be/models"
	"personal-web-be/services"
	"personal-web-be/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)



func LoginHandler(ctx *fiber.Ctx) error {
	var reqLogin model.LoginModel

	
	if err := ctx.BodyParser(&reqLogin); err != nil{
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "invalid Request Body", 
		})
	}

	if err := utils.Validate.Struct(reqLogin); err != nil {
		errors := make(map[string]string)
		for _, v := range err.(validator.ValidationErrors) {
			switch v.Field() {
			case "Email" :
				if v.Tag() == "required" {
					errors["email"] = "Email is Required"
				}else if v.Tag() == "email" {
					errors["email"] = "Email is not valid"
				}
			case "Password" : 
			if v.Tag() == "required" {
				errors["password"] = "Password is Required"
			}else if v.Tag() == "min" {
				errors["passowrd"] = "Password Minium 6 Character"
			}
			}
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : errors,
		})
	}

	token, err := services.LoginServices(reqLogin)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error" : "Email/Password is Wrong" ,
		})
	}


	return ctx.JSON(fiber.Map{
		"message" : "Success Login",
		"token" : token,	
	})
}
