package work

import (
	model "personal-web-be/models"
	"personal-web-be/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)




func GetWorkHandler(ctx *fiber.Ctx) error  {
	data, err := services.GetWorkService()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error" : "Cannot Get Data" + err.Error(), 
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "Success get Data",
		"data" : data,
	})
}


func CreateWorkHandler(ctx *fiber.Ctx)  error {
	uuidGen := uuid.New()
	var work model.WorkModel

	if err := ctx.BodyParser(&work); err != nil{
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "invalid Request Body", 
		})
	}
	work.Id = uuidGen.String()

	err := services.CreateWorkService(work)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error" : err.Error(),
		})
	}


	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "Success Create Work Experience",
		"data" : work,
	})

}

func DeleteWorkHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot Parse UUID with Error"+ err.Error(),
		})
	}

	err = services.DeleteWorkService(uuid)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error Delete with Error"+ err.Error(),
		})
	}


	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "Success Delete Stack",
	})
}