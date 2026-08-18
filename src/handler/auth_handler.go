package handler

import (
	"fmt"
	"time"

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
	payload := new(payload.SignUpPayload)

	if err := c.Bind().Body(payload); err != nil {
		return response.ErrorResponse(c, nil, nil)
	}

	res, err := this.Service.SignUp(payload)

	if err != nil {
		return err.ToResponse(c)
	}
	c.Cookie(&fiber.Cookie{Name: "access_token", Value: res.AccessToken, Path: "/", Expires: time.Now().Add(24 * time.Hour), SameSite: "lax", HTTPOnly: true})
	c.Cookie(&fiber.Cookie{Name: "refresh_token", Value: res.RefreshToken, Path: "/", Expires: time.Now().Add(7 * 24 * time.Hour), SameSite: "lax", HTTPOnly: true})
	return response.SuccessResponse(c, nil, res, nil)
}

func (this *AuthHandler) SignIn(c fiber.Ctx) error {
	payload := new(payload.SignInPayload)

	if err := c.Bind().Body(payload); err != nil {
		return response.ErrorResponse(c, nil, nil)
	}

	res, err := this.Service.SignIn(payload)

	if err != nil {
		return err.ToResponse(c)
	}
	c.Cookie(&fiber.Cookie{Name: "access_token", Value: res.AccessToken, Path: "/", Expires: time.Now().Add(24 * time.Hour), SameSite: "lax", HTTPOnly: true})
	c.Cookie(&fiber.Cookie{Name: "refresh_token", Value: res.RefreshToken, Path: "/", Expires: time.Now().Add(7 * 24 * time.Hour), SameSite: "lax", HTTPOnly: true})
	return response.SuccessResponse(c, nil, res, nil)
}

func (this *AuthHandler) GetUser(c fiber.Ctx) error {
	token := c.Get(fiber.HeaderAuthorization)

	if token == "" {
		status := 401
		message := "Unauthorized"
		fmt.Println("tidak ada token")
		return response.ErrorResponse(c, &message, &status)
	}

	res, err := this.Service.GetUser(token)

	if err != nil {
		return err.ToResponse(c)
	}
	return response.SuccessResponse(c, nil, res, nil)
}
