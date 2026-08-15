package service

import (
	"encoding/json"
	"fmt"

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

func (this *AuthService) SignUp(payload *payload.RegisterPayload) error {
	payload.Data.Role = "user"
	stringBody, err := this.SupabaseService.AuthSignUp(payload)

	fmt.Println(stringBody)

	if err != nil {
		return err
	}

	successResult := new(response.AuthSuccessResult)

	err = json.Unmarshal([]byte(stringBody), successResult)

	if err != nil {
		return err
	}

	return nil
}
