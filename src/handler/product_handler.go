package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/error/service_error"
	"github.com/lalatina11/markita.git/src/lib/payload"
	"github.com/lalatina11/markita.git/src/lib/response"
	"github.com/lalatina11/markita.git/src/model"
	"github.com/lalatina11/markita.git/src/service"
)

type ProductHandler struct {
	ProductService *service.ProductService
}

func NewProductHandler(ProductService *service.ProductService) *ProductHandler {
	return &ProductHandler{ProductService}
}

func (this *ProductHandler) GetAllProducts(c fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	perPageQuery := c.Query("perPage", "")
	if perPageQuery == "" {
		perPageQuery = c.Query("per_page", "25")
	}

	perPage, err := strconv.Atoi(perPageQuery)
	if err != nil || perPage < 1 {
		perPage = 25
	}

	paginatedProducts, serviceErr := this.ProductService.GetAllProducts(page, perPage)
	if serviceErr != nil {
		return serviceErr.ToResponse(c)
	}

	return response.SuccessResponse(c, nil, paginatedProducts, nil)
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

func (this *ProductHandler) PutCategory(c fiber.Ctx) error {
	payload := new(payload.AssignCategoryProductPayload)
	UserID := fiber.Locals[string](c, "user_id")

	if err := c.Bind().Body(payload); err != nil {
		return service_error.Create(422, err.Error()).ToResponse(c)
	}

	updatedProduct, err := this.ProductService.PutCategory(payload, UserID)

	if err != nil {
		return err.ToResponse(c)
	}

	return response.SuccessResponse(c, nil, updatedProduct, nil)
}
