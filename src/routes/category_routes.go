package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/handler"
	"github.com/lalatina11/markita.git/src/service"
	"gorm.io/gorm"
)

func CategoryRoutes(api fiber.Router, Db *gorm.DB) *fiber.Router {
	r := api.Group("/category")

	service := service.NewCategoryService(Db)
	handler := handler.NewCategoryHandler(service)

	r.Get("/", handler.GetALl)

	return &r

}
