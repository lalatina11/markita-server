package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/handler"
	"github.com/lalatina11/markita.git/src/middleware"
)

func StoreRoutes(api fiber.Router) *fiber.Router {
	r := api.Group("/store")

	handler := handler.NewStoreHandler()

	r.Post("/", middleware.AuthMiddleware(), handler.CreateStore)
	r.Get("/:id", handler.Find)
	r.Get("/", handler.GetAll)

	return &r
}
