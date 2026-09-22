package service

import (
	"math"

	"github.com/google/uuid"
	"github.com/lalatina11/markita.git/src/dto/product_dto"
	"github.com/lalatina11/markita.git/src/error/service_error"
	"github.com/lalatina11/markita.git/src/model"
	"gorm.io/gorm"
)

type ProductService struct {
	Db           *gorm.DB
	StoreService *StoreService
}

func NewProductService(Db *gorm.DB) *ProductService {
	StoreService := NewStoreService(Db)
	return &ProductService{Db, StoreService}
}

func (this *ProductService) GetAllProducts(page int, perPage int) (*product_dto.PaginatedProductsDTO, *service_error.ServiceError) {
	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 25
	}

	var total int64
	err := this.Db.Model(&model.Product{}).Count(&total).Error
	if err != nil {
		return nil, service_error.InternalServerError()
	}

	offset := (page - 1) * perPage
	totalPages := 0
	if total > 0 {
		totalPages = int(math.Ceil(float64(total) / float64(perPage)))
	}

	products := []model.Product{}
	err = this.Db.
		Preload("Store").
		Preload("Media").
		Preload("Categories").
		Offset(offset).
		Limit(perPage).
		Order("created_at DESC").
		Find(&products).Error

	if err != nil {
		return nil, service_error.InternalServerError()
	}

	var _products = make([]product_dto.ProductWithRelations, len(products))
	for i, product := range products {
		_products[i] = *product.ToProductDTO()
	}

	return &product_dto.PaginatedProductsDTO{
		Products:   _products,
		Page:       page,
		PerPage:    perPage,
		Total:      total,
		TotalPages: totalPages,
	}, nil
}

func (this *ProductService) CreateProduct(product *model.Product, user_id string) (*product_dto.ProductWithRelations, *service_error.ServiceError) {
	store, err := this.StoreService.Find(product.StoreID)
	if err != nil {
		return nil, service_error.Create(500, "Invalid store ID")
	}

	if store.OwnerID != user_id {
		return nil, service_error.Forbidden()
	}

	productId := uuid.NewString()
	product.ID = productId

	for i := range product.Media {
		product.Media[i].ID = uuid.NewString()
		product.Media[i].ProductID = product.ID
	}

	insertProductErr := this.Db.Create(product).Error

	if insertProductErr != nil {
		return nil, service_error.InternalServerError()
	}

	_product, findProductErr := this.Find(product.ID)

	if findProductErr != nil {
		return nil, service_error.InternalServerError()
	}

	return _product, nil
}

func (this *ProductService) Find(id string) (*product_dto.ProductWithRelations, *service_error.ServiceError) {
	product := new(model.Product)
	product.ID = id

	err := this.Db.
		Preload("Store").
		Preload("Media").
		Preload("Categories").
		First(product).Error

	if err != nil {
		return nil, service_error.NotFound()
	}

	return product.ToProductDTO(), nil
}
