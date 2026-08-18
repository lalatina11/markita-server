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

type SupabaseService struct {
	Config *config.SupabaseConfig
}

func NewSupabaseService() *SupabaseService {
	Config := config.NewSupabaseConfig()
	return &SupabaseService{Config}
}

func (this *SupabaseService) AuthSignUp(payload *payload.SignUpPayload) (string, error) {
	jsonData, err := json.Marshal(payload)

	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/signup", this.Config.AuthURL)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))

	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	apikey := fmt.Sprintf("Bearer %s", this.Config.PublishableKey)
	req.Header.Set("apikey", apikey)

	client := &http.Client{}

	// Send the request
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	stringBody := string(body)
	return stringBody, nil
}

func (this *SupabaseService) AuthSignIn(payload *payload.SignInPayload) (string, error) {
	jsonData, err := json.Marshal(payload)

	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/token?grant_type=password", this.Config.AuthURL)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))

	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	apikey := fmt.Sprintf("Bearer %s", this.Config.PublishableKey)
	req.Header.Set("apikey", apikey)

	client := &http.Client{}

	// Send the request
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	stringBody := string(body)
	return stringBody, nil
}

func (this *SupabaseService) AuthGetUser(token string) (string, error) {

	url := fmt.Sprintf("%s/token?grant_type=password", this.Config.AuthURL)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	apikey := fmt.Sprintf("Bearer %s", this.Config.PublishableKey)
	req.Header.Set("apikey", apikey)
	req.Header.Set("Authorization", token)

	client := &http.Client{}

	// Send the request
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	stringBody := string(body)
	return stringBody, nil
}
