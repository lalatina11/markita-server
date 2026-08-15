package routes

import "github.com/gofiber/fiber/v3"

func AuthRoutes(api *fiber.App) *fiber.Router {
	routes := api.Group("/auth")

	return &routes
}
