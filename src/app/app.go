package app

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/config"
	"github.com/lalatina11/markita.git/src/model"
	"github.com/lalatina11/markita.git/src/routes"
	"github.com/lalatina11/markita.git/src/service"
)

func App() *fiber.App {
	db := config.NewDatabaseConfig().Connect()
	db.AutoMigrate(&model.User{}, &model.Store{}, &model.Product{}, &model.ProductMedia{}, &model.Cart{}, &model.Order{}, &model.OrderItem{}, &model.UserAddress{}, &model.Category{})
	app := fiber.New(fiber.Config{
		BodyLimit: 500 * 1024 * 1024, // 500 MB
	})

	service.NewCategoryService(db).RunSeeder()

	routes.AppRoutes(app, db)

	return app
}
