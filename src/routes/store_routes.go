package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/handler"
	"github.com/lalatina11/markita.git/src/middleware"
	"github.com/lalatina11/markita.git/src/service"
	"gorm.io/gorm"
)

func StoreRoutes(api fiber.Router, Db *gorm.DB) *fiber.Router {
	r := api.Group("/store")

	service := service.NewStoreService(Db)
	handler := handler.NewStoreHandler(service)

	r.Post("/", middleware.AuthMiddleware(Db), handler.CreateStore)
	r.Get("/:id", handler.Find)
	r.Get("/", handler.GetAll)

	return &r
}
