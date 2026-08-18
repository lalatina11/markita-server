package app

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/config"
	"github.com/lalatina11/markita.git/src/model"
	"github.com/lalatina11/markita.git/src/routes"
)

func App() *fiber.App {
	db := config.NewDatabaseConfig().Connect()
	db.AutoMigrate(&model.User{})
	app := fiber.New()

	routes.AppRoutes(app)

	return app
}
