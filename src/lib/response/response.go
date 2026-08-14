package response

import "github.com/gofiber/fiber/v3"

type ApiResponse[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

// SuccessResponse sends a standard JSON success response.
func SuccessResponse[T any](c fiber.Ctx, message *string, data T, statusCode int) error {
	if statusCode == 0 {
		statusCode = fiber.StatusOK
	}

	msg := "OK"
	if message != nil && *message != "" {
		msg = *message
	}

	return c.Status(statusCode).JSON(ApiResponse[T]{
		Success: true,
		Message: msg,
		Data:    data,
	})
}

// ErrorResponse sends a standard JSON error response.
func ErrorResponse(c fiber.Ctx, message *string, statusCode int) error {
	if statusCode == 0 {
		statusCode = fiber.StatusInternalServerError
	}

	msg := "Internal Server Error"
	if message != nil && *message != "" {
		msg = *message
	}

	return c.Status(statusCode).JSON(ApiResponse[any]{
		Success: false,
		Message: msg,
		Data:    nil,
	})
}
