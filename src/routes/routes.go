package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/handler"
)

func AppRoutes(app *fiber.App) {
	appHandler := handler.NewAppHandler()
	app.Get("/", appHandler.Root)
}
