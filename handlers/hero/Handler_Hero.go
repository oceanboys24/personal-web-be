package hero

import (
	"personal-web-be/config"

	"github.com/gofiber/fiber/v2"
)



func  HeroHandler(ctx *fiber.Ctx) error  {
	//var hero model.HeroModel
	var result string

	err := config.SupaClient.DB.From("hero").Select("surname").Single().Execute(&result)
	if err != nil {
		return err
	}


	return ctx.JSON(fiber.Map{
		"data" : result,
	})
}