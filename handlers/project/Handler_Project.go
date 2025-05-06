package project

import (
	model "personal-web-be/models"
	"personal-web-be/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)


func GetProjectHandler(ctx *fiber.Ctx) error  {
	data, err := services.GetProjectService()
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

func GetProjectHandlerById(ctx *fiber.Ctx) error  {
	id := ctx.Params("id")
	data, err := services.GetProjectServiceById(id)
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



func CreateProjectHandler(ctx *fiber.Ctx)  error {
	uuidGen := uuid.New()
	var project model.ProjectModel

	if err := ctx.BodyParser(&project); err != nil{
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "invalid Request Body", 
		})
	}
	project.Id = uuidGen.String()

	err := services.CreateProjectService(project)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error" : err.Error(),
		})
	}


	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "Success Create Project",
		"data" : project,
	})

}

func UpdateProjectHandler(ctx *fiber.Ctx) error   {
	id := ctx.Params("id")
	var project model.ProjectModel

	if err := ctx.BodyParser(&project); err != nil{
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

	if project.Name != "" {
		updatedBody["name"] = project.Name
	}
	if project.Description != "" {
		updatedBody["description"] = project.Description
	}
	if project.Image != "" {
		updatedBody["image_url"] = project.Image
	}
	if project.Repo != nil && *project.Repo != "" {
		updatedBody["repo"] = *project.Repo
	}
	if project.Demo != nil && *project.Demo != "" {
		updatedBody["demo"] = *project.Demo
	}
	if len(project.Stack) >= 0 {
		updatedBody["stack"] = project.Stack
	}

	delete(updatedBody, "id")
	
	err = services.UpdateProjectService(uuid, updatedBody)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error Update with Error"+ err.Error(),
		})
	}


	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "Success Update Project",
	})
}



func DeleteProjectHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot Parse UUID with Error"+ err.Error(),
		})
	}

	err = services.DeleteProjectService(uuid)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error Delete with Error"+ err.Error(),
		})
	}


	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "Success Delete Project",
	})
}