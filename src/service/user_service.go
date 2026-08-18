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
	newUser.Role = payload.User.Role
	newUser.Avatar = avatar
	err := this.Db.Create(newUser).Error

	if err != nil {
		return nil, service_error.InternalServerError()
	}

	return newUser, nil
}

func (this *UserService) FindOrCreate(payload *response.AuthSuccessPayload) (*model.User, *service_error.ServiceError) {
	avatar := fmt.Sprintf("%s%s", config.NewAppConfig().AvatarBaseURL, payload.User.DisplayName)

	user := new(model.User)
	user.ID = payload.User.ID
	user.DisplayName = payload.User.DisplayName
	user.Email = payload.User.Email
	user.Role = payload.User.Role
	err := this.Db.Model(&model.User{}).First(user).Error
	if err != nil {
		user.Avatar = avatar
		_user, err := this.CreateUser(payload)
		if err != nil {
			return nil, service_error.InternalServerError()
		}
		return _user, nil
	}

	return user, nil
}

func (this *UserService) Find(id string) (*model.User, *service_error.ServiceError) {

	user := new(model.User)
	user.ID = id

	err := this.Db.Model(&model.User{}).First(user).Error
	if err != nil {
		return nil, service_error.InternalServerError()
	}

	return user, nil
}
