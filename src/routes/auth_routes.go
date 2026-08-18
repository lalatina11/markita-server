package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/handler"
)

func AuthRoutes(api fiber.Router) *fiber.Router {
	r := api.Group("/auth")

	handler := handler.NewAuthHandler()

	r.Post("/sign-up", handler.SignUp)
	r.Post("/sign-in", handler.SignIn)
	r.Post("/me", handler.GetUser)

	return &r
}
