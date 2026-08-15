package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/lalatina11/markita.git/src/handler"
	"github.com/lalatina11/markita.git/src/middleware"
)

func AppRoutes(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		Format: "# ${method} ${path} - ${status}\n",
	}))

	app.Use(middleware.ErrorHandler())

	appHandler := handler.NewAppHandler()
	app.Get("/", appHandler.Root)
	ApiRoutes(app)

	// ← 404 handler must be LAST, after all routes
	app.Use(middleware.NotFoundHandler())

}
