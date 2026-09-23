package middleware

import (
	"slices"

	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/constants"
	"github.com/lalatina11/markita.git/src/dto/auth_dto"
	"github.com/lalatina11/markita.git/src/error/service_error"
)

func AdminMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		user := fiber.Locals[*auth_dto.AuthUserDto](c, "user")

		if !slices.Contains(constants.ADMIN_ROLES, user.Role) {
			return service_error.Forbidden().ToResponse(c)
		}

		return c.Next()
	}
}
