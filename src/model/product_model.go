package model

import (
	"time"

	"github.com/lalatina11/markita.git/src/dto/product_dto"
	"gorm.io/gorm"
)

type Product struct {
	ID          string         `json:"id" gorm:"primaryKey;index"`
	StoreID     string         `json:"store_id" gorm:"not null;index"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description" gorm:"not null"`
	Price       uint64         `json:"price" gorm:"precission"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoCreateTime;autoUpdateTime:milli"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Store       Store          `json:"store" gorm:"foreignKey:StoreID;references:ID;constraint:OnDelete:CASCADE"`
	Media       []ProductMedia `json:"media" gorm:"foreignKey:ProductID;references:ID"`
}

func (this *Product) ToProductDTO() *product_dto.ProductWithRelations {
	var productMedia = make([]product_dto.ProductMedia, len(this.Media))
	for i, m := range this.Media {
		productMedia[i] = *m.ToProductMediaDTO()
	}
	return &product_dto.ProductWithRelations{
		ID:          this.ID,
		StoreID:     this.StoreID,
		Name:        this.Name,
		Description: this.Description,
		Price:       this.Price,
		CreatedAt:   this.CreatedAt,
		UpdatedAt:   this.UpdatedAt,
		Store: product_dto.ProductStore{
			ID:      this.Store.ID,
			Name:    this.Store.Name,
			OwnerID: this.Store.OwnerID,
			Avatar:  this.Store.Avatar,
			Banner:  this.Store.Banner,
			Address: this.Store.Address,
			City:    this.Store.City,
		},
		Media: productMedia,
	}
}
