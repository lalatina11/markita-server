package model

import (
	"time"

	"gorm.io/gorm"
)

type Store struct {
	ID        string         `json:"id" gorm:"primaryKey;index"`
	Name      string         `json:"name" gorm:"not null;index"`
	OwnerID   string         `json:"owner_id" gorm:"not null;index"`
	Owner     User           `json:"owner" gorm:"foreignKey:OwnerID;references:ID;constraint:OnDelete:CASCADE"`
	Avatar    string         `json:"avatar" gorm:"not null"`
	Banner    string         `json:"banner" gorm:"not null"`
	Address   string         `json:"address" gorm:"not null"`
	City      string         `json:"city" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoCreateTime;autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Products  []Product      `json:"products" gorm:"foreignKey:StoreID;references:ID"`
}
