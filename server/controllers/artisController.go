package controllers

import (
	"server/config"
	"server/helpers"
	"server/models"

	"github.com/gofiber/fiber/v2"
)

func CreateArtis(c *fiber.Ctx) error {
	artis := new(models.Artis)

	if err := c.BodyParser(artis); err != nil {
		return err
	}

	error := helpers.ValidateStruct(c, *artis)
	if error != nil {
		return helpers.Response(c, 400, error, nil)
	}

	if err := config.DB.Create(&artis).Error; err != nil {
		return helpers.Response(c, 500, err.Error(), nil)

	}

	return helpers.Response(c, 201, "Succes Create", nil)
}
