package service

import (
	"encoding/json"

	"github.com/lalatina11/markita.git/src/error/service_error"
	"github.com/lalatina11/markita.git/src/lib/payload"
	"github.com/lalatina11/markita.git/src/lib/response"
)

type AuthService struct {
	SupabaseService *SupabaseService
}

func NewAuthService() *AuthService {
	SupabaseService := NewSupabaseService()
	return &AuthService{SupabaseService}
}

func (this *AuthService) SignUp(payload *payload.RegisterPayload) (*response.AuthSuccessPayload, *service_error.ServiceError) {
	serviceError := service_error.NewServiceError()
	payload.Data.Role = "user"
	stringBody, err := this.SupabaseService.AuthSignUp(payload)

	if err != nil {
		return nil, serviceError
	}

	successResult := new(response.AuthSuccessResult)
	errorResult := new(response.AuthErrorResult)

	err = json.Unmarshal([]byte(stringBody), successResult)

	if err != nil || successResult.AccessToken == "" {
		err = json.Unmarshal([]byte(stringBody), errorResult)

		if err != nil {
			return nil, serviceError
		} else {
			serviceError.Code = errorResult.Code
			serviceError.Msg = errorResult.Msg
			return nil, serviceError
		}
	} else {
		return successResult.ToPayload(), nil
	}

}
