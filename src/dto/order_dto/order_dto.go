package order_dto

import (
	"time"
)

type OrderUser struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	Role        string `json:"role"`
}

type OrderStore struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Avatar  string `json:"avatar"`
	Banner  string `json:"banner"`
	Address string `json:"address"`
	City    string `json:"city"`
	OwnerID string `json:"owner_id"`
}

type OrderProductMedia struct {
	ID        string `json:"id"`
	MediaType string `json:"media_type"`
	MediaURL  string `json:"media_url"`
}

type OrderProduct struct {
	ID          string              `json:"id"`
	StoreID     string              `json:"store_id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Price       uint64              `json:"price"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
	Store       OrderStore          `json:"store"`
	Media       []OrderProductMedia `json:"media"`
}

type OrderItemDTO struct {
	ID             string       `json:"id"`
	OrderID        string       `json:"order_id"`
	ProductID      string       `json:"product_id"`
	Quantity       uint         `json:"quantity"`
	ShippingNumber string       `json:"shipping_number"`
	Status         string       `json:"status"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	Product        OrderProduct `json:"product"`
}

type OrderDestination struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	ReceiverName string `json:"receiver_name"`
	Phone        string `json:"phone"`
	Street       string `json:"street"`
	Description  string `json:"description"`
	City         string `json:"city"`
	Regency      string `json:"regency"`
	Region       string `json:"region"`
}

type OrderDTO struct {
	ID            string           `json:"id"`
	UserID        string           `json:"user_id"`
	DestinationID string           `json:"destination_id"`
	Destination   OrderDestination `json:"destination"`
	CreatedAt     time.Time        `json:"created_at"`
	UpdatedAt     time.Time        `json:"updated_at"`
	User          OrderUser        `json:"user"`
	Items         []OrderItemDTO   `json:"items"`
}
