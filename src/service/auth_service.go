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
	payload.Data.Role = "user"
	stringBody, err := this.SupabaseService.AuthSignUp(payload)
	if err != nil {
		return nil, service_error.NewServiceError()
	}

	var successResult response.AuthSuccessResult
	if err := json.Unmarshal([]byte(stringBody), &successResult); err == nil && successResult.AccessToken != "" {
		return successResult.ToPayload(), nil
	}

	var errorResult response.AuthErrorResult
	if err := json.Unmarshal([]byte(stringBody), &errorResult); err == nil && errorResult.Msg != "" {
		return nil, service_error.CreateServiceError(errorResult.Code, errorResult.Msg)
	}

	return nil, service_error.NewServiceError()
}
