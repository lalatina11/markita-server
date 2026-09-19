package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/handler"
	"github.com/lalatina11/markita.git/src/middleware"
)

func CartRoutes(api fiber.Router) *fiber.Router {
	r := api.Group("/cart")
	handler := handler.NewCartHandler()

	r.Get("/", middleware.AuthMiddleware(), handler.GetAllCarts)
	r.Post("/", middleware.AuthMiddleware(), handler.AddToCart)

	return &r
}
