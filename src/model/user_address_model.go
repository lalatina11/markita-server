package model

import (
	"time"

	"github.com/lalatina11/markita.git/src/dto/auth_dto"
	"gorm.io/gorm"
)

type UserAddress struct {
	ID           string         `json:"id" gorm:"primaryKey;index"`
	UserID       string         `json:"user_id" gorm:"not null"`
	ReceiverName string         `json:"receiver_name" gorm:"not null"`
	Phone        string         `json:"phone" gorm:"not null"`
	Street       string         `json:"street" gorm:"not null"`
	Description  string         `json:"description"`
	City         string         `json:"city" gorm:"not null"`
	Regency      string         `json:"regency" gorm:"not null"`
	Region       string         `json:"region" gorm:"not null"`
	IsDefault    bool           `json:"is_default" gorm:"default:false"`
	CreatedAt    time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"autoCreateTime;autoUpdateTime:milli"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	User         User           `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

func (this *UserAddress) ToUserAddressDTO() *auth_dto.UserAddress {
	return &auth_dto.UserAddress{
		ID:           this.ID,
		UserID:       this.UserID,
		ReceiverName: this.ReceiverName,
		Phone:        this.Phone,
		Street:       this.Street,
		Description:  this.Description,
		City:         this.City,
		Regency:      this.Regency,
		Region:       this.Region,
		IsDefault:    this.IsDefault,
		CreatedAt:    this.CreatedAt,
		UpdatedAt:    this.UpdatedAt,
	}
}
