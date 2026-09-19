package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/handler"
	"github.com/lalatina11/markita.git/src/middleware"
	"github.com/lalatina11/markita.git/src/service"
	"gorm.io/gorm"
)

func AuthRoutes(api fiber.Router, Db *gorm.DB) *fiber.Router {
	r := api.Group("/auth")

	service := service.NewAuthService(Db)
	handler := handler.NewAuthHandler(service)

	r.Post("/sign-up", handler.SignUp)
	r.Post("/sign-in", handler.SignIn)
	r.Post("/refresh", handler.RefreshToken)

	// Protected Routes
	r.Get("/me", middleware.AuthMiddleware(Db), handler.GetUser)
	r.Delete("/sign-out", middleware.AuthMiddleware(Db), handler.SignOut)
	return &r
}
