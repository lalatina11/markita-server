package model

import (
	"time"

	"github.com/lalatina11/markita.git/src/dto/order_dto"
	"gorm.io/gorm"
)

type OrderItem struct {
	ID             string         `json:"id" gorm:"primaryKey;index"`
	OrderID        string         `json:"order_id" gorm:"not null;index:uidx_order_item_order_product,unique"`
	ProductID      string         `json:"product_id" gorm:"not null;index:uidx_order_item_order_product,unique"`
	Quantity       uint           `json:"quantity"`
	ShippingNumber string         `json:"shipping_number" gorm:"not null"`
	Status         string         `json:"status" gorm:"not null"`
	CreatedAt      time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `json:"updated_at" gorm:"autoCreateTime;autoUpdateTime:milli"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Order          Order          `json:"order" gorm:"foreignKey:OrderID;references:ID;constraint:OnDelete:CASCADE"`
	Product        Product        `json:"product" gorm:"foreignKey:ProductID;references:ID;constraint:OnDelete:CASCADE"`
}

func (this *OrderItem) ToOrderItemDTO() *order_dto.OrderItemDTO {
	this.Product.Store.FixMediaURL()

	var media = make([]order_dto.OrderProductMedia, len(this.Product.Media))
	for i, m := range this.Product.Media {
		m.FixMediaURL()
		media[i] = order_dto.OrderProductMedia{
			ID:        m.ID,
			MediaType: m.MediaType,
			MediaURL:  m.MediaURL,
		}
	}

	return &order_dto.OrderItemDTO{
		ID:             this.ID,
		OrderID:        this.OrderID,
		ProductID:      this.ProductID,
		Quantity:       this.Quantity,
		ShippingNumber: this.ShippingNumber,
		Status:         this.Status,
		CreatedAt:      this.CreatedAt,
		UpdatedAt:      this.UpdatedAt,
		Product: order_dto.OrderProduct{
			ID:          this.Product.ID,
			StoreID:     this.Product.StoreID,
			Name:        this.Product.Name,
			Description: this.Product.Description,
			Price:       this.Product.Price,
			CreatedAt:   this.Product.CreatedAt,
			UpdatedAt:   this.Product.UpdatedAt,
			Store: order_dto.OrderStore{
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
