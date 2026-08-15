package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/lib/payload"
	"github.com/lalatina11/markita.git/src/lib/response"
	"github.com/lalatina11/markita.git/src/service"
)

type AuthHandler struct {
	Service *service.AuthService
}

func NewAuthHandler() *AuthHandler {
	Service := service.NewAuthService()
	return &AuthHandler{Service}
}

func (this *AuthHandler) SignUp(c fiber.Ctx) error {
	payload := new(payload.RegisterPayload)

	if err := c.Bind().Body(payload); err != nil {
		return response.ErrorResponse(c, nil, nil)
	}

	res, err := this.Service.SignUp(payload)

	if err != nil {
		return err.ToResponse(c)
	}

	return response.SuccessResponse(c, nil, res, nil)
}
