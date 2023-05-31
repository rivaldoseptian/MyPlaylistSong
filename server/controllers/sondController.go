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

func GetSong(c *fiber.Ctx) error {
	var song []models.Song
	var songResponse []models.SongResponse

	if err := config.DB.Preload("Artis").Find(&song).Find(&songResponse).Error; err != nil {
		return helpers.Response(c, 500, err.Error(), nil)
	}
	return helpers.Response(c, 200, "Succes", songResponse)
}

func GetOneSong(c *fiber.Ctx) error {
	var song models.Song
	var songResponse models.SongResponse

	songId := c.Params("id")

	if err := config.DB.Preload("Artis").First(&song, "id= ?", songId).First(&songResponse).Error; err != nil {
		return helpers.Response(c, 404, "Song Not Found", nil)
	}

	return helpers.Response(c, 200, "Succces", songResponse)

}

func EditSong(c *fiber.Ctx) error {
	songId := c.Params("id")
	var song models.Song

	songRequest := new(models.SongRequest)

	if err := c.BodyParser(songRequest); err != nil {
		return err
	}

	error := helpers.ValidateStruct(c, *songRequest)
	if error != nil {
		return helpers.Response(c, 404, error, nil)
	}

	if err := config.DB.First(&song, "id= ?", songId).Error; err != nil {
		return helpers.Response(c, 404, "Song Not Found", nil)
	}

	song = models.Song{
		Name:     songRequest.Name,
		Duration: songRequest.Duration,
		ArtisID:  songRequest.ArtisID,
	}

	if err := config.DB.Where("id= ?", songId).Updates(&song).Error; err != nil {
		return helpers.Response(c, 500, err.Error(), nil)
	}

	return helpers.Response(c, 200, "Succes Edit Song", nil)
}

func DeleteSong(c *fiber.Ctx) error {
	songId := c.Params("id")
	var song models.Song

	if err := config.DB.First(&song, "id= ?", songId).Error; err != nil {
		return helpers.Response(c, 404, "Song Not Found", nil)
	}

	if err := config.DB.Where("id= ?", songId).Delete(&song).Error; err != nil {
		return helpers.Response(c, 500, err.Error(), nil)
	}

	return helpers.Response(c, 200, "Succes Delete Song", nil)
}
