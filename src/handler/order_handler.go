package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/lib/response"
	"github.com/lalatina11/markita.git/src/model"
	"github.com/lalatina11/markita.git/src/service"
)

type OrderHandler struct {
	OrderService *service.OrderService
}

func NewOrderHandler(OrderService *service.OrderService) *OrderHandler {
	return &OrderHandler{OrderService}
}

func (this *OrderHandler) GetAllOrders(c fiber.Ctx) error {
	userID := fiber.Locals[string](c, "user_id")
	orders, err := this.OrderService.GetAllOrders(userID)

	if err != nil {
		return err.ToResponse(c)
	}

	return response.SuccessResponse(c, nil, orders, nil)
}

func (this *OrderHandler) Direct(c fiber.Ctx) error {
	msg := "Order data created!"
	code := 201

	payload := new(model.Order)
	userID := fiber.Locals[string](c, "user_id")

	if err := c.Bind().Body(payload); err != nil {
		msg = err.Error()
		code = 422
		return response.ErrorResponse(c, &msg, &code)
	}

	order, err := this.OrderService.Direct(payload, userID)

	if err != nil {
		return err.ToResponse(c)
	}

	return response.SuccessResponse(c, &msg, order, &code)

}
