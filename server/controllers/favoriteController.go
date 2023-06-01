package controllers

import (
	"server/config"
	"server/helpers"
	"server/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateFavorite(c *fiber.Ctx) error {
	user := c.Context().UserValue("userinfo").(*helpers.MyCustomClaims)
	songId := c.Params("id")
	id, _ := strconv.Atoi(songId)

	favorite := models.Favorite{
		UserID: user.ID,
		SongID: uint(id),
	}

	if err := config.DB.Create(&favorite).Error; err != nil {
		return helpers.Response(c, 500, err.Error(), nil)
	}

	return helpers.Response(c, 201, "SUcces Create Favorite", nil)

}
