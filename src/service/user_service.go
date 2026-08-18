package service

import (
	"fmt"

	"github.com/lalatina11/markita.git/src/config"
	"github.com/lalatina11/markita.git/src/error/service_error"
	"github.com/lalatina11/markita.git/src/lib/response"
	"github.com/lalatina11/markita.git/src/model"
	"gorm.io/gorm"
)

type UserService struct {
	Db *gorm.DB
}

func NewUserService() *UserService {
	Db := config.NewDatabaseConfig().Connect()
	return &UserService{Db}
}

func (this *UserService) CreateUser(payload *response.AuthSuccessPayload) (*model.User, *service_error.ServiceError) {
	if payload == nil {
		return nil, service_error.Create(422, "Invalid payload")
	}

	avatar := fmt.Sprintf("%s%s", config.NewAppConfig().AvatarBaseURL, payload.User.DisplayName)

	newUser := new(model.User)
	newUser.ID = payload.User.ID
	newUser.DisplayName = payload.User.DisplayName
	newUser.Email = payload.User.Email
	newUser.Avatar = avatar
	err := this.Db.Create(&newUser).Error

	if err != nil {
		return nil, service_error.InternalServerError()
	}

	return newUser, nil
}
