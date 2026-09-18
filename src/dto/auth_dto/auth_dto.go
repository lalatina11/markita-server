package auth_dto

import (
	"time"

	"gorm.io/gorm"
)

type UserStore struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	OwnerID string `json:"owner_id"`
	Avatar  string `json:"avatar"`
	Banner  string `json:"banner"`
	Address string `json:"address"`
	City    string `json:"city"`
}

type UserWithStoresDto struct {
	ID          string         `json:"id"`
	DisplayName string         `json:"display_name"`
	Email       string         `json:"email"`
	Avatar      string         `json:"avatar"`
	Role        string         `json:"role"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	Stores      []UserStore    `json:"stores"`
}
