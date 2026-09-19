package cart_dto

import (
	"time"
)

type CartUser struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	Role        string `json:"role"`
}

type CartProduct struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	Role        string `json:"role"`
}

type CartDTO struct {
	ID        string    `json:"id" gorm:"primaryKey;index"`
	UserID    string    `json:"user_id" gorm:"not null;index:uidx_cart_user_product,unique"`
	ProductID string    `json:"product_id" gorm:"not null;index:uidx_cart_user_product,unique"`
	Quantity  uint      `json:"quantity"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime;autoUpdateTime:milli"`
	User      CartUser  `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}
