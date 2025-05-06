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


func GetWorkHandlerById(ctx *fiber.Ctx) error  {
	id := ctx.Params("id")
	data, err := services.GetWorkServiceById(id)
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
func UpdateWorkHandler(ctx *fiber.Ctx) error   {
	id := ctx.Params("id")
	var work model.WorkModel

	if err := ctx.BodyParser(&work); err != nil{
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "invalid Request Body" + err.Error(), 
		})
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot Parse UUID with Error"+ err.Error(),
		})
	}

	updatedBody := make(map[string]interface{})

	if work.Role != "" {
		updatedBody["role"] = work.Role
	}
	if work.Company != "" {
		updatedBody["company"] = work.Company
	}
	if work.Image != "" {
		updatedBody["image_url"] = work.Image
	}
	if work.StartDate != nil && *work.StartDate != "" {
		updatedBody["start_date"] = *work.StartDate
	}
	if work.EndDate != nil && *work.EndDate != "" {
		updatedBody["end_date"] = *work.EndDate
	}
	if len(work.Task) > 0 {
		updatedBody["task"] = work.Task
	}
	if len(work.Stack) > 0 {
		updatedBody["stack"] = work.Stack
	}


	delete(updatedBody, "id")
	
	err = services.UpdateWorkService(uuid, updatedBody)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error Update with Error"+ err.Error(),
		})
	}


	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "Success Update Work Experience",
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
		"message" : "Success Delete Work",
	})
}