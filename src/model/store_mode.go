package model

import (
	"time"

	"gorm.io/gorm"
)

type Store struct {
	ID        string         `json:"id" gorm:"primaryKey,index"`
	Name      string         `json:"name" gorm:"not null,index"`
	OwnerID   string         `json:"owner_id" gorm:"not null,index"`
	Owner     User           `json:"owner" gorm:"foreignKey:OwnerID;references:ID;constraint:OnDelete:CASCADE"`
	Address   string         `json:"address" gorm:"not null"`
	City      string         `json:"city" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
