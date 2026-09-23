package model

import (
	"time"

	"github.com/lalatina11/markita.git/src/dto/category_dto"
	"gorm.io/gorm"
)

type Category struct {
	ID        string         `json:"id" gorm:"primaryKey;index"`
	Name      string         `json:"name" gorm:"not null;index"`
	Slug      string         `json:"slug" gorm:"not null;unique;index"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoCreateTime;autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Products  []Product      `json:"products" gorm:"many2many:product_categories"`
}

func (this *Category) ToCategoryDTO() *category_dto.CategoryDTO {
	return &category_dto.CategoryDTO{
		ID:        this.ID,
		Name:      this.Name,
		Slug:      this.Slug,
		CreatedAt: this.CreatedAt,
		UpdatedAt: this.UpdatedAt,
	}
}
