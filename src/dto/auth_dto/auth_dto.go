package auth_dto

import (
	"time"
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

type UserAddress struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	ReceiverName string    `json:"receiver_name"`
	Phone        string    `json:"phone"`
	Street       string    `json:"street"`
	Description  string    `json:"description"`
	City         string    `json:"city"`
	Regency      string    `json:"regency"`
	Region       string    `json:"region"`
	IsDefault    bool      `json:"is_default"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserWithStoresDto struct {
	ID          string        `json:"id"`
	DisplayName string        `json:"display_name"`
	Email       string        `json:"email"`
	Avatar      string        `json:"avatar"`
	Role        string        `json:"role"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	Stores      []UserStore   `json:"stores"`
	Addresses   []UserAddress `json:"addresses"`
}
