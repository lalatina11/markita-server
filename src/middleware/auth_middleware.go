package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/error/service_error"
	"github.com/lalatina11/markita.git/src/service"
	"gorm.io/gorm"
)

func AuthMiddleware(Db *gorm.DB) fiber.Handler {
	authService := service.NewAuthService(Db)
	return func(c fiber.Ctx) error {
		token := c.Get(fiber.HeaderAuthorization)

		user, err := authService.GetUser(token)

		if err != nil {
			return service_error.Unauthorized().ToResponse(c)
		}

		c.Locals("token", token)
		c.Locals("user_id", user.ID)
		c.Locals("user", user)
		return c.Next()
	}
}
