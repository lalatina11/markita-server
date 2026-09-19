package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/handler"
	"github.com/lalatina11/markita.git/src/middleware"
	"github.com/lalatina11/markita.git/src/service"
	"gorm.io/gorm"
)

func CartRoutes(api fiber.Router, Db *gorm.DB) *fiber.Router {
	r := api.Group("/cart")

	service := service.NewCartService(Db)
	handler := handler.NewCartHandler(service)

	r.Get("/", middleware.AuthMiddleware(Db), handler.GetAllCarts)
	r.Post("/", middleware.AuthMiddleware(Db), handler.AddToCart)

	return &r
}
