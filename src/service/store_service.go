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

	updateUserErr := this.UserService.UpdateRole(userId, "seller")

	if updateUserErr != nil {
		return nil, updateUserErr
	}

	return store, nil
}

func (this *StoreService) Find(id string) (*model.Store, *service_error.ServiceError) {
	store := new(model.Store)
	store.ID = id

	err := this.Db.Preload("Owner").Model(&model.Store{}).First(store).Error

	if err != nil {
		return nil, service_error.NotFound()
	}

	return store, nil
}

func (this *StoreService) GetAll() ([]model.Store, *service_error.ServiceError) {
	stores := []model.Store{}

	err := this.Db.Preload("Owner").Model(&model.Store{}).Find(&stores).Error

	if err != nil {
		return nil, service_error.InternalServerError()
	}

	return stores, nil

}
