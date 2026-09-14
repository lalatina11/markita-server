package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/handler"
	"github.com/lalatina11/markita.git/src/middleware"
)

func AuthRoutes(api fiber.Router) *fiber.Router {
	r := api.Group("/auth")

	handler := handler.NewAuthHandler()

	r.Post("/sign-up", handler.SignUp)
	r.Post("/sign-in", handler.SignIn)

	// Protected Routes
	r.Get("/me", middleware.AuthMiddleware(), handler.GetUser)
	r.Delete("/sign-out", middleware.AuthMiddleware(), handler.SignOut)
	return &r
}
