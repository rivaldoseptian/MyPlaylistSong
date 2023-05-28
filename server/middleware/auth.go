package middleware

import (
	"server/helpers"

	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	accessToken := c.Get("accesstoken")

	if accessToken == "" {
		return helpers.Response(c, 401, "Unauthorized", nil)
	}

	user, err := helpers.ValidateToken(accessToken)

	if err != nil {
		return helpers.Response(c, 401, "Unauthorized", nil)
	}

	c.Locals("userinfo", user)

	return c.Next()
}
