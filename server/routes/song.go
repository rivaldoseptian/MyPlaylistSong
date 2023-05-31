package routes

import (
	"server/controllers"
	"server/middleware"

	"github.com/gofiber/fiber/v2"
)

func SongRouter(r *fiber.App) {
	adminRouter := r.Group("/admin")
	adminRouter.Use(middleware.Auth)
	adminRouter.Use(middleware.AdminAuth)
	adminRouter.Post("/song", controllers.CreateSong)

}
