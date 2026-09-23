package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/lib/response"
	"gorm.io/gorm"
)

func ApiRoutes(app *fiber.App, Db *gorm.DB) *fiber.Router {
	api := app.Group("/api")
	api.Get("/", func(c fiber.Ctx) error {
		return response.SuccessResponse[any](c, nil, nil, nil)
	})

	AuthRoutes(api, Db)
	UploadRoutes(api)
	StoreRoutes(api, Db)
	ProductRoutes(api, Db)
	CartRoutes(api, Db)
	OrderRoutes(api, Db)
	UserAddressRoutes(api, Db)
	CategoryRoutes(api, Db)

	return &api
}
