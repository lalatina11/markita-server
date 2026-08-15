// middleware/error_handler.go
package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/lib/response"
)

type AppError struct {
	Message    string
	StatusCode int
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(message string, statusCode int) *AppError {
	return &AppError{Message: message, StatusCode: statusCode}
}

func ErrorHandler() fiber.Handler {
	return func(c fiber.Ctx) error {
		if err := c.Next(); err != nil {
			return response.ErrorResponse(c, nil, nil)
		}
		return nil
	}
}

func NotFoundHandler() fiber.Handler {
	return func(c fiber.Ctx) error {
		status := 404
		msg := "Not Found Route!"
		return response.ErrorResponse(c, &msg, &status)
	}
}
