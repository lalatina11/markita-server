package service

import (
	"github.com/google/uuid"
	"github.com/lalatina11/markita.git/src/config"
	"github.com/lalatina11/markita.git/src/error/service_error"
	"github.com/lalatina11/markita.git/src/model"
	"gorm.io/gorm"
)

type StoreService struct {
	Db          *gorm.DB
	UserService *UserService
}

func NewStoreService() *StoreService {
	Db := config.NewDatabaseConfig().Connect()
	UserService := NewUserService()
	return &StoreService{Db, UserService}
}

func (this *StoreService) CreateStore(store *model.Store, userId string) (*model.Store, *service_error.ServiceError) {
	store.ID = uuid.NewString()
	store.OwnerID = userId
	err := this.Db.Create(store).Error
	if err != nil {
		return nil, service_error.CreateServiceError(500, "Failed to create a store")
	}

	user, _err := this.UserService.Find(userId)
	if _err != nil {
		return nil, _err
	}

	updateUserErr := this.UserService.UpdateRole(user.ID, "seller")

	if updateUserErr != nil {
		return nil, updateUserErr
	}

	return store, nil
}
