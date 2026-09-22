package model

import (
	"time"

	"github.com/lalatina11/markita.git/src/dto/order_dto"
	"gorm.io/gorm"
)

type Order struct {
	ID          string         `json:"id" gorm:"primaryKey;index"`
	UserID      string         `json:"user_id" gorm:"not null"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoCreateTime;autoUpdateTime:milli"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	User        User           `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Items       []OrderItem    `json:"items" gorm:"foreignKey:OrderID;references:ID;constraint:OnDelete:CASCADE"`
	AddressID   string         `json:"destination_id" gorm:"not null"`
	Destination UserAdress     `json:"destination" gorm:"foreignKey:AddressID;references:ID"`
}

func (this *Order) ToOrderDTO() *order_dto.OrderDTO {
	var items = make([]order_dto.OrderItemDTO, len(this.Items))
	for i, item := range this.Items {
		items[i] = *item.ToOrderItemDTO()
	}

	return &order_dto.OrderDTO{
		ID:        this.ID,
		UserID:    this.UserID,
		CreatedAt: this.CreatedAt,
		UpdatedAt: this.UpdatedAt,
		User: order_dto.OrderUser{
			ID:          this.User.ID,
			DisplayName: this.User.DisplayName,
			Email:       this.User.Email,
			Avatar:      this.User.Avatar,
			Role:        this.User.Role,
		},
		Items: items,
	}
}
