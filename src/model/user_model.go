package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          string `json:"id" gorm:"primaryKey,index"`
	DisplayName string `json:"display_name" gorm:"index"`
	Email       string `json:"email" gorm:"unique,index"`
	Avatar      string `json:"avatar"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
