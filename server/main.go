package main

import (
	"fmt"
	"os"
	"server/config"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config.LoadEnv()
	config.ConnectDB()

	app := fiber.New()

	//routes
	routes.IndexRouter(app)

	app.Listen(fmt.Sprintf(":%v", os.Getenv("PORT")))
}
