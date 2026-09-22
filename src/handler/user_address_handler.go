package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/lib/response"
	"github.com/lalatina11/markita.git/src/model"
	"github.com/lalatina11/markita.git/src/service"
)

type UserAddressHandler struct {
	UserAddressService *service.UserAddressService
}

func NewUserAddressHandler(UserAddressService *service.UserAddressService) *UserAddressHandler {
	return &UserAddressHandler{UserAddressService}
}

func (this *UserAddressHandler) Create(c fiber.Ctx) error {
	msg := "Address added successfully"
	code := 201
	UserID := fiber.Locals[string](c, "user_id")
	payload := new(model.UserAddress)

	if err := c.Bind().Body(payload); err != nil {
		msg = err.Error()
		code = 422
		return response.ErrorResponse(c, &msg, &code)
	}

	address, err := this.UserAddressService.Create(payload, UserID)

	if err != nil {
		return err.ToResponse(c)
	}

	return response.SuccessResponse(c, &msg, address, &code)

}
