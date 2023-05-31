package controllers

import (
	"server/config"
	"server/helpers"
	"server/models"

	"github.com/gofiber/fiber/v2"
)

func CreateSong(c *fiber.Ctx) error {
	reqSong := new(models.SongRequest)

	if err := c.BodyParser(reqSong); err != nil {
		return err
	}

	error := helpers.ValidateStruct(c, *reqSong)
	if error != nil {
		return helpers.Response(c, 400, error, nil)
	}

	song := models.Song{
		Name:     reqSong.Name,
		Duration: reqSong.Duration,
		ArtisID:  reqSong.ArtisID,
	}

	if err := config.DB.Create(&song).Error; err != nil {
		return helpers.Response(c, 500, err.Error(), nil)
	}

	return helpers.Response(c, 201, "Succes Create", nil)
}
