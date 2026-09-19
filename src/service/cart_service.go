package service

import (
	"github.com/google/uuid"
	"github.com/lalatina11/markita.git/src/config"
	"github.com/lalatina11/markita.git/src/dto/cart_dto"
	"github.com/lalatina11/markita.git/src/error/service_error"
	"github.com/lalatina11/markita.git/src/model"
	"gorm.io/gorm"
)

type CartService struct {
	Db *gorm.DB
}

func NewCartSerice() *CartService {
	Db := config.NewDatabaseConfig().Connect()
	return &CartService{Db}
}

func (this *CartService) Find(id string) (*cart_dto.CartDTO, *service_error.ServiceError) {
	cart := new(model.Cart)
	cart.ID = id

	err := this.Db.
		Preload("User").
		Preload("Product").
		Preload("Product.Store").
		Preload("Product.Media").
		First(cart).Error

	if err != nil {
		return nil, service_error.NotFound()
	}

	return cart.ToCartDTO(), nil
}

func (this *CartService) FindByProductIdAndUserId(productID string, userID string) (*model.Cart, *service_error.ServiceError) {
	cart := new(model.Cart)
	err := this.Db.Where(&model.Cart{ProductID: productID, UserID: userID}).First(cart).Error

	if err != nil {
		return nil, service_error.NotFound()
	}

	return cart, nil
}

func (this *CartService) AddToCart(cart *model.Cart, user_id string) (*cart_dto.CartDTO, *service_error.ServiceError) {
	existingCart, _ := this.FindByProductIdAndUserId(cart.ProductID, user_id)

	existingProduct, serviceErr := NewProductService().Find(cart.ProductID)

	if serviceErr != nil {
		return nil, service_error.Create(400, "Invalid Product")
	}

	if existingProduct.Store.OwnerID == user_id {
		return nil, service_error.Create(400, "You can not add your product into your cart!")
	}

	if existingCart != nil {
		existingCart.Quantity = existingCart.Quantity + cart.Quantity
		if err := this.Db.Save(existingCart).Error; err != nil {
			return nil, service_error.InternalServerError()
		}
		return this.Find(existingCart.ID)
	}

	cart.ID = uuid.NewString()
	cart.UserID = user_id
	err := this.Db.Create(cart).Error
	if err != nil {
		return nil, service_error.Create(500, "Failed to Add to cart!")
	}

	return this.Find(cart.ID)
}

func (this *CartService) GetAllCarts(userID string) ([]cart_dto.CartDTO, *service_error.ServiceError) {
	var carts []model.Cart

	err := this.Db.
		Preload("User").
		Preload("Product").
		Preload("Product.Store").
		Preload("Product.Media").
		Where(&model.Cart{UserID: userID}).
		Find(&carts).Error

	if err != nil {
		return nil, service_error.InternalServerError()
	}

	var dtos = make([]cart_dto.CartDTO, len(carts))
	for i, c := range carts {
		dtos[i] = *c.ToCartDTO()
	}

	return dtos, nil
}
