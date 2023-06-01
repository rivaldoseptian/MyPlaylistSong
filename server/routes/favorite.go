package routes

import (
	"server/controllers"
	"server/middleware"

	"github.com/gofiber/fiber/v2"
)

func FavoriteRouter(r *fiber.App) {
	userRouter := r.Group("/user")
	userRouter.Use(middleware.Auth)
	userRouter.Get("/favorite", controllers.GetFavoriteSong)
	userRouter.Post("/favorite/:id", controllers.CreateFavorite)
	userRouter.Delete("/favorite/:id", controllers.DeleteFavoriteSong)

}
