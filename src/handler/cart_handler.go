package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/lib/response"
	"github.com/lalatina11/markita.git/src/model"
	"github.com/lalatina11/markita.git/src/service"
)

type CartHandler struct {
	CartService *service.CartService
}

func NewCartHandler() *CartHandler {
	CartService := service.NewCartSerice()
	return &CartHandler{CartService}
}

func (this *CartHandler) AddToCart(c fiber.Ctx) error {
	msg := "Success to add to cart!"
	code := 201
	payload := new(model.Cart)
	userId := fiber.Locals[string](c, "user_id")

	if err := c.Bind().Body(payload); err != nil {
		msg = err.Error()
		code = 422
		return response.ErrorResponse(c, &msg, &code)
	}

	cart, serviceErr := this.CartService.AddToCart(payload, userId)

	if serviceErr != nil {
		return serviceErr.ToResponse(c)
	}

	return response.SuccessResponse(c, &msg, cart, &code)
}

func (this *CartHandler) GetAllCarts(c fiber.Ctx) error {
	userId := fiber.Locals[string](c, "user_id")

	carts, serviceErr := this.CartService.GetAllCarts(userId)

	if serviceErr != nil {
		return serviceErr.ToResponse(c)
	}

	return response.SuccessResponse(c, nil, carts, nil)
}
