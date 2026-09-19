package product_dto

import (
	"time"
)

type ProductMedia struct {
	ID        string `json:"id"`
	MediaType string `json:"media_type"`
	MediaURL  string `json:"media_url"`
}

type ProductStore struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	OwnerID string `json:"owner_id"`
	Avatar  string `json:"avatar"`
	Banner  string `json:"banner"`
	Address string `json:"address"`
	City    string `json:"city"`
}

type ProductWithRelations struct {
	ID          string         `json:"id"`
	StoreID     string         `json:"store_id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       uint64         `json:"price"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Store       ProductStore   `json:"store"`
	Media       []ProductMedia `json:"media"`
}

type PaginatedProductsDTO struct {
	Products   []ProductWithRelations `json:"products"`
	Page       int                    `json:"page"`
	PerPage    int                    `json:"per_page"`
	Total      int64                  `json:"total"`
	TotalPages int                    `json:"total_pages"`
}
