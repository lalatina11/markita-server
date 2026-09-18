package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/lib/response"
	"github.com/lalatina11/markita.git/src/model"
	"github.com/lalatina11/markita.git/src/service"
)

type StoreHandler struct {
	StoreService *service.StoreService
}

func NewStoreHandler() *StoreHandler {
	StoreService := service.NewStoreService()
	return &StoreHandler{StoreService}
}

func (this *StoreHandler) CreateStore(c fiber.Ctx) error {
	payload := new(model.Store)
	code := 201
	msg := "Success to create a store!"
	if err := c.Bind().Body(payload); err != nil {
		msg = err.Error()
		code = 422
		return response.ErrorResponse(c, &msg, &code)
	}
	userId := fiber.Locals[string](c, "user_id")
	res, err := this.StoreService.CreateStore(payload, userId)
	if err != nil {
		return err.ToResponse(c)
	}
	return response.SuccessResponse(c, &msg, res, &code)
}

func (this *StoreHandler) Find(c fiber.Ctx) error {
	id := c.Params("id")

	store, err := this.StoreService.Find(id)
	if err != nil {
		return err.ToResponse(c)
	}

	return response.SuccessResponse(c, nil, store, nil)
}

func (this *StoreHandler) GetAll(c fiber.Ctx) error {
	stores, err := this.StoreService.GetAll()
	if err != nil {
		return err.ToResponse(c)
	}

	return response.SuccessResponse(c, nil, stores, nil)
}
