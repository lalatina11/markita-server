package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/handler"
	"github.com/lalatina11/markita.git/src/middleware"
)

func ProductRoutes(api fiber.Router) *fiber.Router {
	r := api.Group("/product")
	handler := handler.NewProductHandler()

	r.Get("/", handler.GetAllProducts)
	r.Get("/:id", handler.Find)

	r.Post("/", middleware.AuthMiddleware(), handler.CreateProduct)

	return &r

}
