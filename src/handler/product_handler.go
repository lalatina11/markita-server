package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/lib/response"
	"github.com/lalatina11/markita.git/src/model"
	"github.com/lalatina11/markita.git/src/service"
)

type ProductHandler struct {
	ProductService *service.ProductService
}

func NewProductHandler() *ProductHandler {
	ProductService := service.NewProductService()
	return &ProductHandler{ProductService}
}

func (this *ProductHandler) GetAllProducts(c fiber.Ctx) error {
	products, err := this.ProductService.GetAllProducts()
	if err != nil {
		return err.ToResponse(c)
	}

	return response.SuccessResponse(c, nil, products, nil)
}

func (this *ProductHandler) CreateProduct(c fiber.Ctx) error {
	payload := new(model.Product)
	msg := "Success to create product!"
	code := 201

	if err := c.Bind().Body(payload); err != nil {
		msg = err.Error()
		code = 422
		return response.ErrorResponse(c, &msg, &code)
	}

	userId := fiber.Locals[string](c, "user_id")
	product, err := this.ProductService.CreateProduct(payload, userId)

	if err != nil {
		return err.ToResponse(c)
	}

	return response.SuccessResponse(c, nil, product, nil)
}

func (this *ProductHandler) Find(c fiber.Ctx) error {
	id := c.Params("id")

	product, err := this.ProductService.Find(id)

	if err != nil {
		return err.ToResponse(c)
	}

	return response.SuccessResponse(c, nil, product, nil)

}
