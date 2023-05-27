package routes

import (
	"server/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(r *fiber.App) {
	adminRouter := r.Group("/admin")
	adminRouter.Post("/register", controllers.RegisterAdmin)
}
