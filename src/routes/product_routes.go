package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/handler"
	"github.com/lalatina11/markita.git/src/middleware"
	"github.com/lalatina11/markita.git/src/service"
	"gorm.io/gorm"
)

func ProductRoutes(api fiber.Router, Db *gorm.DB) *fiber.Router {
	r := api.Group("/product")

	service := service.NewProductService(Db)
	handler := handler.NewProductHandler(service)

	r.Put("/product-category", middleware.AuthMiddleware(Db), handler.PutCategory)

	r.Get("/", handler.GetAllProducts)
	r.Post("/", middleware.AuthMiddleware(Db), handler.CreateProduct)

	r.Get("/:id", handler.Find)

	return &r
}
