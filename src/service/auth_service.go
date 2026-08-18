package service

import (
	"encoding/json"

	"github.com/lalatina11/markita.git/src/error/service_error"
	"github.com/lalatina11/markita.git/src/lib/payload"
	"github.com/lalatina11/markita.git/src/lib/response"
	"github.com/lalatina11/markita.git/src/lib/validator"
)

type AuthService struct {
	SupabaseService *SupabaseService
	UserService     *UserService
}

func NewAuthService() *AuthService {
	SupabaseService := NewSupabaseService()
	UserService := NewUserService()
	return &AuthService{SupabaseService, UserService}
}

func (this *AuthService) SignUp(payload *payload.SignUpPayload) (*response.AuthUserPayload, *service_error.ServiceError) {
	payload.Data.Role = "user"
	errs := validator.Validate(payload)
	if errs != nil {
		return nil, errs[0].ToServiceError()
	}
	stringBody, err := this.SupabaseService.AuthSignUp(payload)
	if err != nil {
		return nil, service_error.NewServiceError()
	}

	var successResult response.AuthSuccessResult
	if err := json.Unmarshal([]byte(stringBody), &successResult); err == nil && successResult.IsSuccess() {
		payload := successResult.ToPayload()
		user, err := this.UserService.CreateUser(payload)
		if err != nil {
			return nil, service_error.Create(500, "Failed to create User")
		}
		return payload.ToAuthUserPayload(user), nil
	}

	var errorResult response.AuthErrorResult
	if err := json.Unmarshal([]byte(stringBody), &errorResult); err == nil && errorResult.Msg != "" {
		return nil, service_error.Create(errorResult.Code, errorResult.Msg)
	}

	return nil, service_error.NewServiceError()
}

func (this *AuthService) SignIn(payload *payload.SignInPayload) (*response.AuthUserPayload, *service_error.ServiceError) {
	errs := validator.Validate(payload)
	if errs != nil {
		return nil, errs[0].ToServiceError()
	}
	var successResult response.AuthSuccessResult
	stringBody, err := this.SupabaseService.AuthSignIn(payload)
	if err != nil {
		return nil, service_error.NewServiceError()
	}
	if err := json.Unmarshal([]byte(stringBody), &successResult); err == nil && successResult.IsSuccess() {
		payload := successResult.ToPayload()
		user, err := this.UserService.CreateUser(payload)
		if err != nil {
			return nil, service_error.Create(500, "Failed to create User")
		}
		return payload.ToAuthUserPayload(user), nil
	}

	var errorResult response.AuthErrorResult
	if err := json.Unmarshal([]byte(stringBody), &errorResult); err == nil && errorResult.Msg != "" {
		return nil, service_error.Create(errorResult.Code, errorResult.Msg)
	}
	return nil, service_error.NewServiceError()
}
