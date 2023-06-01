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

func GetFavoriteSong(c *fiber.Ctx) error {
	var favorite []models.Favorite
	var favoriteResponse []models.FavoriteResponse
	user := c.Context().UserValue("userinfo").(*helpers.MyCustomClaims)

	if err := config.DB.Where("user_id", user.ID).Preload("Song.Artis").Find(&favorite).Find(&favoriteResponse).Error; err != nil {
		return helpers.Response(c, 500, err.Error(), nil)
	}

	return helpers.Response(c, 200, "Success", favoriteResponse)
}

func DeleteFavoriteSong(c *fiber.Ctx) error {
	favoriteId := c.Params("id")
	var favorite models.Favorite

	if err := config.DB.First(&favorite, "id= ?", favoriteId).Error; err != nil {
		return helpers.Response(c, 404, "Favorite Song Not Found", nil)
	}

	if err := config.DB.Where("id= ?", favoriteId).Delete(&favorite).Error; err != nil {
		return helpers.Response(c, 500, err.Error(), nil)
	}

	return helpers.Response(c, 200, "Succes Delete Favorite Song", nil)

}
