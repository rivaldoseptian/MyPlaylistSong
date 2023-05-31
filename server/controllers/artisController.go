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

func GetArtis(c *fiber.Ctx) error {
	var artis []models.Artis

	if err := config.DB.Find(&artis).Error; err != nil {
		return helpers.Response(c, 500, err.Error(), nil)
	}

	return helpers.Response(c, 200, "Succes", artis)
}

func GetOneArtis(c *fiber.Ctx) error {
	var artis models.Artis

	artisId := c.Params("id")

	if err := config.DB.First(&artis, "id = ?", artisId).Error; err != nil {
		return helpers.Response(c, 404, "User Not Found", nil)
	}

	return helpers.Response(c, 200, "Success", artis)
}
