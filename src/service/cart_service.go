package service

import (
	"github.com/google/uuid"
	"github.com/lalatina11/markita.git/src/config"
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

func (this *CartService) FindByProductIdAndUserId(productID string, userID string) (*model.Cart, *service_error.ServiceError) {
	cart := new(model.Cart)
	cart.ProductID = productID
	cart.UserID = userID
	err := this.Db.First(cart).Error

	if err != nil {
		return nil, service_error.NotFound()
	}

	return cart, nil
}

func (this *CartService) AddToCart(cart *model.Cart, user_id string) (*model.Cart, *service_error.ServiceError) {
	existingCart, _ := this.FindByProductIdAndUserId(cart.ProductID, user_id)

	if existingCart != nil {
		existingCart.Quantity = existingCart.Quantity + cart.Quantity
		this.Db.Save(existingCart)
		return existingCart, nil
	}

	cart.ID = uuid.NewString()
	cart.UserID = user_id
	err := this.Db.Create(cart).Error
	if err != nil {
		return nil, service_error.Create(500, "Failed to Add to cart!")
	}
	return cart, nil

}

func (this *CartService) GetAllCarts(userID string) ([]model.Cart, *service_error.ServiceError) {
	var carts []model.Cart

	err := this.Db.Where(&model.Cart{UserID: userID}).Find(&carts).Error

	if err != nil {
		return nil, service_error.InternalServerError()
	}

	return carts, nil
}
