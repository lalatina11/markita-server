package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/handler"
	"github.com/lalatina11/markita.git/src/middleware"
)

func OrderRoutes(api fiber.Router) *fiber.Router {
	r := api.Group("/order")

	handler := handler.NewOrderHandler()

	r.Get("/", middleware.AuthMiddleware(), handler.GetAllOrders)
	r.Post("/direct", middleware.AuthMiddleware(), handler.Direct)

	return &r
}
