package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/lib/response"
	"github.com/lalatina11/markita.git/src/service"
)

type CategoryHandler struct {
	CategoryService *service.CategoryService
}

func NewCategoryHandler(CategoryService *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{CategoryService}
}

func (this *CategoryHandler) GetALl(c fiber.Ctx) error {
	categories, err := this.CategoryService.GetAll()
	if err != nil {
		return err.ToResponse(c)
	}

	return response.SuccessResponse(c, nil, categories, nil)
}
