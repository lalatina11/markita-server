package model

import (
	"time"

	"gorm.io/gorm"
)

type UserAdress struct {
	ID          string         `json:"id" gorm:"primaryKey;index"`
	UserID      string         `json:"user_id" gorm:"not null"`
	Street      string         `json:"street" gorm:"not null"`
	Description string         `json:"description"`
	City        string         `json:"city" gorm:"not null"`
	Regency     string         `json:"regency" gorm:"not null"`
	Region      string         `json:"region" gorm:"not null"`
	Phone       string         `json:"phone" gorm:"not null"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoCreateTime;autoUpdateTime:milli"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	User        User           `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}
