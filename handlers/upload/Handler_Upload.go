package upload

import (
	"personal-web-be/utils"

	"github.com/gofiber/fiber/v2"
)



func UploadHandler(c *fiber.Ctx) error {
	file ,err := c.FormFile("image")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message" : "cannot read file",
		})
	}


	srcFile , err := file.Open()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message" : "cannot open file",
		})
	}	
	
	resultDownload , err := utils.UploadFile("personal-web",file.Filename, srcFile, file.Size, file.Header.Get("content-type"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message " : "Upload Failed",
		})
	}


	return c.Status(200).JSON(fiber.Map{
		"message " : "Success Upload",
		"link" : resultDownload,
	})
}