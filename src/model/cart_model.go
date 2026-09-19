package model

import (
	"time"

	"gorm.io/gorm"
)

type CartModel struct {
	ID        string         `json:"id" gorm:"primaryKey;index"`
	UserID    string         `json:"user_id" gorm:"not null"`
	ProductID string         `json:"product_id" gorm:"not null"`
	Quantity  uint           `json:"quantity"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoCreateTime;autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	User      User           `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Product   Product        `json:"product" gorm:"foreignKey:ProductID;references:ID;constraint:OnDelete:CASCADE"`
}
