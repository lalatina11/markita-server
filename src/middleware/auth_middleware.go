package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/error/service_error"
	"github.com/lalatina11/markita.git/src/service"
)

func AuthMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		token := c.Get(fiber.HeaderAuthorization)

		user, err := service.NewAuthService().GetUser(token)

		if err != nil {
			return service_error.Unauthorized().ToResponse(c)
		}

		c.Locals("token", token)
		c.Locals("user_id", user.ID)
		c.Locals("user", user)
		return c.Next()
	}
}
