package cart_dto

import (
	"time"
)

type CartUser struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	Role        string `json:"role"`
}

type CartStore struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Avatar  string `json:"avatar"`
	Banner  string `json:"banner"`
	Address string `json:"address"`
	City    string `json:"city"`
	OwnerID string `json:"owner_id"`
}

type CartProductMedia struct {
	ID        string `json:"id"`
	MediaType string `json:"media_type"`
	MediaURL  string `json:"media_url"`
}

type CartProduct struct {
	ID          string             `json:"id"`
	StoreID     string             `json:"store_id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Price       uint64             `json:"price"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	Store       CartStore          `json:"store"`
	Media       []CartProductMedia `json:"media"`
}

type CartDTO struct {
	ID        string      `json:"id"`
	UserID    string      `json:"user_id"`
	ProductID string      `json:"product_id"`
	Quantity  uint        `json:"quantity"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	User      CartUser    `json:"user"`
	Product   CartProduct `json:"product"`
}
