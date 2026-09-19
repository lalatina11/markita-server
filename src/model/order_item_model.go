package model

import (
	"time"

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
}
