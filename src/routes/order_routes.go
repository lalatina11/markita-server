package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/handler"
	"github.com/lalatina11/markita.git/src/middleware"
	"github.com/lalatina11/markita.git/src/service"
	"gorm.io/gorm"
)

func OrderRoutes(api fiber.Router, Db *gorm.DB) *fiber.Router {
	r := api.Group("/order")

	service := service.NewOrderService(Db)
	handler := handler.NewOrderHandler(service)

	r.Get("/", middleware.AuthMiddleware(Db), handler.GetAllOrders)
	r.Post("/direct", middleware.AuthMiddleware(Db), handler.Direct)

	return &r
}
