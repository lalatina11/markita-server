package model

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        string         `json:"id" gorm:"primaryKey;index"`
	Name      string         `json:"name" gorm:"not null;index"`
	Slug      string         `json:"slug" gorm:"not null;unique;index"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoCreateTime;autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Products  []Product      `json:"products" gorm:"many2many:product_categories"`
}
