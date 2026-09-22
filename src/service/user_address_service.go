package service

import (
	"github.com/google/uuid"
	"github.com/lalatina11/markita.git/src/error/service_error"
	"github.com/lalatina11/markita.git/src/model"
	"gorm.io/gorm"
)

type UserAddressService struct {
	Db *gorm.DB
}

func NewUserAdressService(Db *gorm.DB) *UserAddressService {
	return &UserAddressService{Db}
}

func (this *UserAddressService) IsFirstAddress(UserID string) bool {
	address := new(model.UserAddress)
	address.UserID = UserID

	err := this.Db.First(address)

	if err != nil {
		return false
	}

	return true
}

func (this *UserAddressService) Create(adress *model.UserAddress, UserID string) (*model.UserAddress, *service_error.ServiceError) {
	adress.ID = uuid.NewString()
	adress.UserID = UserID
	isFirstAddress := this.IsFirstAddress(UserID)

	if isFirstAddress {
		adress.IsPrimary = true
	}

	err := this.Db.Create(adress).Error

	if err != nil {
		return nil, service_error.InternalServerError()
	}

	return adress, nil
}
