package service

import (
	"fmt"
	"slices"
	"time"

	"github.com/lalatina11/markita.git/src/config"
	"github.com/lalatina11/markita.git/src/constants"
	"github.com/lalatina11/markita.git/src/error/service_error"
	supabaseresponse "github.com/lalatina11/markita.git/src/lib/response/supabase_response"
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

func (this *UserService) CreateUser(payload *supabaseresponse.AuthSuccessPayload) (*model.User, *service_error.ServiceError) {
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

	if !slices.Contains(constants.ALLOWED_USER_ROLES, newUser.Role) {
		newUser.Role = "user"
	}

	if err != nil {
		return nil, service_error.InternalServerError()
	}

	return newUser, nil
}

func (this *UserService) FindOrCreate(payload *supabaseresponse.AuthSuccessPayload) (*model.User, *service_error.ServiceError) {
	avatar := fmt.Sprintf("%s%s", config.NewAppConfig().AvatarBaseURL, payload.User.DisplayName)

	user := new(model.User)
	user.ID = payload.User.ID
	err := this.Db.Model(&model.User{}).First(user).Error
	if err != nil {
		user.DisplayName = payload.User.DisplayName
		user.Email = payload.User.Email
		if !slices.Contains(constants.ALLOWED_USER_ROLES, user.Role) {
			user.Role = "user"
		}
		user.Role = payload.User.Role
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

	err := this.Db.Preload("Stores").Model(&model.User{}).First(user).Error
	if err != nil {
		return nil, service_error.InternalServerError()
	}

	return user, nil
}

func (this *UserService) UpdateRole(id string, role string) *service_error.ServiceError {

	if !slices.Contains(constants.ALLOWED_USER_ROLES, role) {
		return service_error.CreateServiceError(422, "Invalid role!")
	}

	user, err := this.Find(id)

	if err != nil {
		return service_error.NotFound()
	}

	user.ID = id
	user.Role = role
	user.UpdatedAt = time.Now()

	_err := this.Db.Save(user).Error
	if _err != nil {
		return service_error.CreateServiceError(500, "Failed to update user role!")
	}
	return nil
}
