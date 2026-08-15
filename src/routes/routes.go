package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/lalatina11/markita.git/src/handler"
)

func AppRoutes(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		Format: "# ${method} ${path} - ${status}\n",
	}))

	appHandler := handler.NewAppHandler()
	app.Get("/", appHandler.Root)
}
