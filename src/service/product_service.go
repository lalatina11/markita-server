package service

import (
	"github.com/google/uuid"
	"github.com/lalatina11/markita.git/src/config"
	"github.com/lalatina11/markita.git/src/dto/product_dto"
	"github.com/lalatina11/markita.git/src/error/service_error"
	"github.com/lalatina11/markita.git/src/model"
	"gorm.io/gorm"
)

type ProductService struct {
	Db           *gorm.DB
	StoreService *StoreService
}

func NewProductService() *ProductService {
	Db := config.NewDatabaseConfig().Connect()
	StoreService := NewStoreService()
	return &ProductService{Db, StoreService}
}

func (this *ProductService) GetAllProducts() ([]product_dto.ProductWithRelations, *service_error.ServiceError) {
	products := []model.Product{}

	err := this.Db.Preload("Store").Preload("Media").Model(model.Product{}).Find(&products).Error

	if err != nil {
		return nil, service_error.InternalServerError()
	}

	var _products = make([]product_dto.ProductWithRelations, len(products))

	for i, product := range products {
		_products[i] = *product.ToProductDTO()
	}

	return _products, nil
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

	err := this.Db.Preload("Store").Preload("Media").First(product).Error

	if err != nil {
		return nil, service_error.NotFound()
	}

	return product.ToProductDTO(), nil
}
