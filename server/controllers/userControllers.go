package controllers

import (
	"server/config"
	"server/helpers"
	"server/models"

	"github.com/gofiber/fiber/v2"
)

func RegisterAdmin(c *fiber.Ctx) error {
	register := new(models.Register)

	if err := c.BodyParser(register); err != nil {
		return err
	}

	error := helpers.ValidateStruct(c, *register)
	if error != nil {
		return helpers.Response(c, 400, error, nil)
	}

	passwordHash, err := helpers.HassingPassword(register.Password)

	if err != nil {
		return helpers.Response(c, 500, err.Error(), nil)
	}

	user := models.User{
		Name:     register.Name,
		Email:    register.Email,
		Password: passwordHash,
		Role:     "Admin",
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return helpers.Response(c, 500, err.Error(), nil)
	}

	return helpers.Response(c, 201, "Success", user)
}
