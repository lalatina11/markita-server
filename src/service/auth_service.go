package service

import (
	"fmt"

	"github.com/lalatina11/markita.git/src/lib/payload"
)

type AuthService struct {
	SupabaseService *SupabaseService
}

func NewAuthService() *AuthService {
	SupabaseService := NewSupabaseService()
	return &AuthService{SupabaseService}
}

func (this *AuthService) SignUp(payload *payload.RegisterPayload) error {
	stringBody, err := this.SupabaseService.AuthSignUp(payload)

	if err != nil {
		return err
	}

	fmt.Println(stringBody)

	return nil
}
