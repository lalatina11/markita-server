package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/lib/response"
)

type AppHandler struct{}

func NewAppHandler() *AppHandler {
	return &AppHandler{}
}

func (a *AppHandler) Root(c fiber.Ctx) error {
	return response.SuccessResponse[any](c, nil, nil, fiber.StatusOK)
}
