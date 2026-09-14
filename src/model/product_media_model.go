package model

import (
	"time"

	"github.com/lalatina11/markita.git/src/utils"
	"gorm.io/gorm"
)

type ProductMedia struct {
	ID        string         `json:"id" gorm:"primaryKey;index"`
	ProductID string         `json:"product_id" gorm:"not null;index"`
	Product   Product        `json:"product" gorm:"foreignKey:ProductID;references:ID;constraint:OnDelete:CASCADE"`
	MediaType string         `json:"media_type" gorm:"not null"`
	MediaURL  string         `json:"media_url" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoCreateTime;autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (this *ProductMedia) FixMediaURL() *ProductMedia {
	this.MediaURL = utils.NewCommonUtiliity().GenerateMediaURL(this.MediaURL)
	return this
}
