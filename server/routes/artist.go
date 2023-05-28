package routes

import (
	"server/controllers"
	"server/middleware"

	"github.com/gofiber/fiber/v2"
)

func ArtisRouter(r *fiber.App) {
	adminRouter := r.Group("/admin")
	adminRouter.Use(middleware.Auth)
	adminRouter.Use(middleware.AdminAuth)
	adminRouter.Post("/artis", controllers.CreateArtis)
}
