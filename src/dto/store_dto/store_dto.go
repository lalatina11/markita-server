package store_dto

import "time"

type StoreOwner struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	Role        string `json:"role"`
}

type StoreWithOwnerDTO struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	OwnerID   string     `json:"owner_id"`
	Avatar    string     `json:"avatar"`
	Banner    string     `json:"banner"`
	Address   string     `json:"address"`
	City      string     `json:"city"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Owner     StoreOwner `json:"owner"`
}
