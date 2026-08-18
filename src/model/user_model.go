package model

import (
	"time"

	"gorm.io/gorm"
)

type UserRole string

const (
	Regular UserRole = "user"
	Admin   UserRole = "admin"
)

type User struct {
	ID          string         `json:"id" gorm:"primaryKey,index"`
	DisplayName string         `json:"display_name" gorm:"index"`
	Email       string         `json:"email" gorm:"unique,index"`
	Avatar      string         `json:"avatar"`
	Role        UserRole       `json:"role"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
