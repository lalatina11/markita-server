package model

import (
	"time"

	"github.com/lalatina11/markita.git/src/dto/auth_dto"
	"gorm.io/gorm"
)

type User struct {
	ID          string         `json:"id" gorm:"primaryKey;index"`
	DisplayName string         `json:"display_name" gorm:"index"`
	Email       string         `json:"email" gorm:"unique;index"`
	Avatar      string         `json:"avatar"`
	Role        string         `json:"role"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoCreateTime;autoUpdateTime:milli"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Stores      []Store        `json:"stores" gorm:"foreignKey:OwnerID;references:ID"`
}

func (this *User) ToUserWithStoreDTO() *auth_dto.UserWithStoresDto {
	var stores = make([]auth_dto.UserStore, len(this.Stores))

	for i, store := range this.Stores {
		stores[i] = *store.ToUserWithStoreDTO()
	}

	return &auth_dto.UserWithStoresDto{
		ID:          this.ID,
		DisplayName: this.DisplayName,
		Email:       this.Email,
		Role:        this.Role,
		Avatar:      this.Avatar,
		CreatedAt:   this.CreatedAt,
		UpdatedAt:   this.UpdatedAt,
		Stores:      stores,
	}
}
