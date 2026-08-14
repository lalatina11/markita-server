package lib

import "github.com/gofiber/fiber/v3"

type ApiResponse[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

// SuccessResponse sends a standard JSON success response.
func SuccessResponse[T any](c fiber.Ctx, statusCode int, message string, data T) error {
	if statusCode == 0 {
		statusCode = fiber.StatusOK
	}
	return c.Status(statusCode).JSON(ApiResponse[any]{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse sends a standard JSON error response.
func ErrorResponse(c fiber.Ctx, statusCode int, message string) error {
	if statusCode == 0 {
		statusCode = fiber.StatusInternalServerError
	}
	if message == "" {
		message = "Internal Server Error"
	}
	return c.Status(statusCode).JSON(ApiResponse[any]{
		Success: false,
		Message: message,
		Data:    nil,
	})
}
