package service_error

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/lib/response"
)

type ServiceError struct {
	Code int
	Msg  string
}

func NewServiceError() *ServiceError {
	return &ServiceError{Code: 500, Msg: "Internal Server Error"}
}

func CreateServiceError(Code int, Msg string) *ServiceError {
	return &ServiceError{Code, Msg}
}

func Unauthorized() *ServiceError {
	return &ServiceError{Code: 401, Msg: "Unauthorized"}
}

func Forbidden() *ServiceError {
	return &ServiceError{Code: 403, Msg: "Forbidden"}
}

func (this *ServiceError) ToResponse(c fiber.Ctx) error {
	return response.ErrorResponse(c, &this.Msg, &this.Code)
}
