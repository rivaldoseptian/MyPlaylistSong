package routes

import (
	"server/controllers"
	"server/middleware"

	"github.com/gofiber/fiber/v2"
)

func FavoriteRouter(r *fiber.App) {
	userRouter := r.Group("/user")
	userRouter.Use(middleware.Auth)
	userRouter.Post("/favorite/:id", controllers.CreateFavorite)

}
