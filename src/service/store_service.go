package service

import (
	"github.com/google/uuid"
	"github.com/lalatina11/markita.git/src/config"
	"github.com/lalatina11/markita.git/src/dto/store_dto"
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

	return store, nil
}

func (this *StoreService) Find(id string) (*store_dto.StoreWithOwnerDTO, *service_error.ServiceError) {
	store := new(model.Store)
	store.ID = id

	err := this.Db.Preload("Owner").Model(&model.Store{}).First(store).Error

	if err != nil {
		return nil, service_error.NotFound()
	}

	return store.ToStoreDTO(), nil
}

func (this *StoreService) GetAll() ([]store_dto.StoreWithOwnerDTO, *service_error.ServiceError) {
	stores := []model.Store{}

	err := this.Db.Preload("Owner").Model(&model.Store{}).Find(&stores).Error

	if err != nil {
		return nil, service_error.InternalServerError()
	}

	var dto = make([]store_dto.StoreWithOwnerDTO, len(stores))

	for i, store := range stores {
		dto[i] = *store.ToStoreDTO()
	}

	return dto, nil

}
