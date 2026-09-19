package model

import (
	"time"

	"github.com/lalatina11/markita.git/src/dto/cart_dto"
	"gorm.io/gorm"
)

type Cart struct {
	ID        string         `json:"id" gorm:"primaryKey;index"`
	UserID    string         `json:"user_id" gorm:"not null;index:uidx_cart_user_product,unique"`
	ProductID string         `json:"product_id" gorm:"not null;index:uidx_cart_user_product,unique"`
	Quantity  uint           `json:"quantity"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoCreateTime;autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	User      User           `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Product   Product        `json:"product" gorm:"foreignKey:ProductID;references:ID;constraint:OnDelete:CASCADE"`
}

func (this *Cart) ToCartDTO() *cart_dto.CartDTO {
	this.Product.Store.FixMediaURL()

	var media = make([]cart_dto.CartProductMedia, len(this.Product.Media))
	for i, m := range this.Product.Media {
		m.FixMediaURL()
		media[i] = cart_dto.CartProductMedia{
			ID:        m.ID,
			MediaType: m.MediaType,
			MediaURL:  m.MediaURL,
		}
	}

	return &cart_dto.CartDTO{
		ID:        this.ID,
		UserID:    this.UserID,
		ProductID: this.ProductID,
		Quantity:  this.Quantity,
		CreatedAt: this.CreatedAt,
		UpdatedAt: this.UpdatedAt,
		User: cart_dto.CartUser{
			ID:          this.User.ID,
			DisplayName: this.User.DisplayName,
			Email:       this.User.Email,
			Avatar:      this.User.Avatar,
			Role:        this.User.Role,
		},
		Product: cart_dto.CartProduct{
			ID:          this.Product.ID,
			StoreID:     this.Product.StoreID,
			Name:        this.Product.Name,
			Description: this.Product.Description,
			Price:       this.Product.Price,
			CreatedAt:   this.Product.CreatedAt,
			UpdatedAt:   this.Product.UpdatedAt,
			Store: cart_dto.CartStore{
				ID:      this.Product.Store.ID,
				Name:    this.Product.Store.Name,
				Avatar:  this.Product.Store.Avatar,
				Banner:  this.Product.Store.Banner,
				Address: this.Product.Store.Address,
				City:    this.Product.Store.City,
				OwnerID: this.Product.Store.OwnerID,
			},
			Media: media,
		},
	}
}
