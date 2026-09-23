package service

import (
	"github.com/google/uuid"
	"github.com/lalatina11/markita.git/src/dto/category_dto"
	"github.com/lalatina11/markita.git/src/error/service_error"
	"github.com/lalatina11/markita.git/src/model"
	"github.com/lalatina11/markita.git/src/utils"
	"gorm.io/gorm"
)

type CategoryService struct {
	Db *gorm.DB
}

func NewCategoryService(Db *gorm.DB) *CategoryService {
	return &CategoryService{Db}
}

func (this *CategoryService) RunSeeder() *service_error.ServiceError {
	var existingCategoryNumber int64
	err := this.Db.Model(&model.Category{}).Count(&existingCategoryNumber).Error

	if err != nil {
		return service_error.InternalServerError()
	}

	if existingCategoryNumber > 0 {
		return nil
	}

	util := utils.NewCommonUtiliity()
	cateoryNames := []string{"Fashion", "Elektronik", "Sepatu", "Pakaian", "Produk Digital", "Makanan", "Minuman", "F&B", "Alat Rumah Tangga"}
	categories := make([]model.Category, len(cateoryNames))

	for i := range categories {
		categories[i].ID = uuid.NewString()
		name := cateoryNames[i]
		categories[i].Name = name
		categories[i].Slug = util.SlugifyUnicode(name)
	}

	err = this.Db.Create(&categories).Error

	if err != nil {
		return service_error.InternalServerError()
	}

	return nil
}

func (this *CategoryService) GetAll() ([]category_dto.CategoryDTO, *service_error.ServiceError) {
	var categories []model.Category

	err := this.Db.Find(&categories).Error

	if err != nil {
		return nil, service_error.InternalServerError()
	}

	dto := make([]category_dto.CategoryDTO, len(categories))

	for i := range categories {
		dto[i] = *categories[i].ToCategoryDTO()
	}

	return dto, nil
}
