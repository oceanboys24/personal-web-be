package stack

import (
	model "personal-web-be/models"
	"personal-web-be/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetStackHandler(ctx *fiber.Ctx) error  {
	result, err  := services.GetStackService()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error" : "Cannot Get Stack",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "Success Get Stack",
		"data" : result,
	})
}



func CreateStackHandler(ctx *fiber.Ctx)  error {
	uuidGen := uuid.New()
	var stack model.StackModel

	if err := ctx.BodyParser(&stack); err != nil{
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "invalid Request Body", 
		})
	}
	stack.Id = uuidGen.String()
	

	err := services.CreateStackService(stack)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error" : err.Error(),
		})
	}


	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "Success Create Stack",
		"data" : stack,
	})

}

func DeleteStackHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot Parse UUID with Error"+ err.Error(),
		})
	}

	err = services.DeleteStackService(uuid)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error Delete with Error"+ err.Error(),
		})
	}


	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "Success Delete Stack",
	})
}