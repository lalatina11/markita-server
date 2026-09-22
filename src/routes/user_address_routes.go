package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/handler"
	"github.com/lalatina11/markita.git/src/middleware"
	"github.com/lalatina11/markita.git/src/service"
	"gorm.io/gorm"
)

func UserAddressRoutes(api fiber.Router, Db *gorm.DB) *fiber.Router {
	r := api.Group("/user-address")

	service := service.NewUserAdressService(Db)
	handler := handler.NewUserAddressHandler(service)

	r.Post("/", middleware.AuthMiddleware(Db), handler.Create)

	return &r
}
