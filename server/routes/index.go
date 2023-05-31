package routes

import "github.com/gofiber/fiber/v2"

func IndexRouter(r *fiber.App) {
	AuthRouter(r)
	ArtisRouter(r)
	SongRouter(r)
}
