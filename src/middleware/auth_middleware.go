package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/lib/response"
	"github.com/lalatina11/markita.git/src/service"
)

func AuthMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		token := c.Get(fiber.HeaderAuthorization)

		user, err := service.NewAuthService().GetUser(token)

		if err != nil {
			message := "Error from Auth Middleware"
			return response.ErrorResponse(c, &message, nil)
		}

		c.Locals("token", token)
		c.Locals("user_id", user.ID)
		c.Locals("user", user)
		return c.Next()
	}
}
