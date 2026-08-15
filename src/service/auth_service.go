package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/lalatina11/markita.git/src/config"
	"github.com/lalatina11/markita.git/src/lib/payload"
)

type AuthService struct {
	Supabase *config.SupabaseConfig
}

func NewAuthService() *AuthService {
	Supabase := config.NewSupabaseConfig()
	return &AuthService{Supabase}
}

func (this *AuthService) SignUp(payload *payload.RegisterPayload) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	signUpURL := fmt.Sprintf("%s/signup", this.Supabase.AuthURL)

	req, err := http.NewRequest(http.MethodPost, signUpURL, bytes.NewBuffer(jsonData))

	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	authorization := fmt.Sprintf("Bearer %s", this.Supabase.PublishableKey)
	req.Header.Set("Authorization", authorization)

	client := &http.Client{}

	// Send the request
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	stringBody := string(body)

	fmt.Println(stringBody)

	return nil
}
