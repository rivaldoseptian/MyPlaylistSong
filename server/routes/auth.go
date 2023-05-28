package routes

import (
	"server/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(r *fiber.App) {
	adminRouter := r.Group("/admin")
	adminRouter.Post("/register", controllers.RegisterAdmin)
	adminRouter.Post("/login", controllers.LoginUser)

	userRouter := r.Group("/user")
	userRouter.Post("/register", controllers.RegisterUser)
	userRouter.Post("/login", controllers.LoginUser)
}
