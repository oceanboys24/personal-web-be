package hero

import (
	model "personal-web-be/models"
	"personal-web-be/services"

	"github.com/gofiber/fiber/v2"
)

func HeroHandlerGet(ctx *fiber.Ctx)  error {
	result, err := services.GetServiceHero() 
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error" : err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "Success Get Hero",
			"data" : result,
	})	
}



func HeroHandlerUpdate(ctx *fiber.Ctx) error  {
	var body model.HeroModel

	if err := ctx.BodyParser(&body); err != nil{
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "invalid Request Body" +  err.Error(), 
		})
	}

	updatedBody := make(map[string]interface{})

	if body.Surname != "" {
		updatedBody["surname"] = body.Surname
	}
	if body.Profession != "" {
		updatedBody["profession"] = body.Profession
	}
	if body.Email != "" {
		updatedBody["email"] = body.Email
	}
	if body.Cv != "" {
		updatedBody["cv"] = body.Cv
	}
	if body.Handphone != "" {
		updatedBody["handphone"] = body.Handphone
	}

	if body.Description != "" {
		updatedBody["description"] = body.Description
	}
	if body.Location != "" {
		updatedBody["location"] = body.Location
	}

	if body.IsAvailable != nil {
		updatedBody["is_available"] = *body.IsAvailable
	}


	if body.ImageUrl != "" {
		updatedBody["image_url"] = body.ImageUrl
	}


	if len(updatedBody) == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "no valid fields to update",
		})
	}

	result , err := services.UpdateServiceHero(updatedBody)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error" : err.Error(),
		})
	}
	
	return ctx.JSON(fiber.Map{
		"message" : "Success Update",
		"data" : result,
	})
}