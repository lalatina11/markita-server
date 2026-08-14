package app

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/routes"
)

func App() *fiber.App {
	app := fiber.New()

	routes.AppRoutes(app)

	return app
}
