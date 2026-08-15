package response

import "github.com/gofiber/fiber/v3"

type ApiResponse[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

// SuccessResponse sends a standard JSON success response.
func SuccessResponse[T any](c fiber.Ctx, message *string, data T, statusCode *int) error {
	code := 200
	if statusCode != nil && *statusCode != 0 {
		code = *statusCode
	}

	msg := "OK"
	if message != nil && *message != "" {
		msg = *message
	}

	return c.Status(code).JSON(ApiResponse[T]{
		Success: true,
		Message: msg,
		Data:    data,
	})
}

// ErrorResponse sends a standard JSON error response.
func ErrorResponse(c fiber.Ctx, message *string, statusCode *int) error {
	code := 500
	if statusCode != nil && *statusCode != 0 {
		code = *statusCode
	}
	msg := "Internal Server Error"
	if message != nil && *message != "" {
		msg = *message
	}

	return c.Status(code).JSON(ApiResponse[any]{
		Success: false,
		Message: msg,
		Data:    nil,
	})
}
