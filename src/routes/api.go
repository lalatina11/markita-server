package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/lib/response"
)

func ApiRoutes(app *fiber.App) *fiber.Router {
	api := app.Group("/api")
	api.Get("/", func(c fiber.Ctx) error {
		return response.SuccessResponse[any](c, nil, nil, nil)
	})

	return &api
}
