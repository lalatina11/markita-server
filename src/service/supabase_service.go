package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/config"
	"github.com/lalatina11/markita.git/src/lib/payload"
	"github.com/lalatina11/markita.git/src/utils"
)

type SupabaseService struct {
	Config *config.SupabaseConfig
	Util   *utils.CommonUtility
}

func NewSupabaseService() *SupabaseService {
	Config := config.NewSupabaseConfig()
	Util := utils.NewCommonUtiliity()
	return &SupabaseService{Config, Util}
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

	url := fmt.Sprintf("%s/user", this.Config.AuthURL)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	apikey := fmt.Sprintf("Bearer %s", this.Config.PublishableKey)
	req.Header.Set("apikey", apikey)
	req.Header.Set(fiber.HeaderAuthorization, token)

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

func (this *SupabaseService) AuthSignOut(access_token string) error {

	url := fmt.Sprintf("%s/logout", this.Config.AuthURL)

	req, err := http.NewRequest(http.MethodPost, url, nil)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	apikey := fmt.Sprintf("Bearer %s", this.Config.PublishableKey)
	req.Header.Set("apikey", apikey)
	req.Header.Set(fiber.HeaderAuthorization, access_token)

	client := &http.Client{}

	// Send the request
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return nil
}

func (this *SupabaseService) AuthRefreshToken(payload *payload.RefreshTokenPayload) (string, error) {
	jsonData, err := json.Marshal(payload)

	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/token?grant_type=refresh_token", this.Config.AuthURL)

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

func (this *SupabaseService) StorageUploadFile(payload *payload.FileUploadPayload) (string, error) {
	url := this.Config.StorageURL
	file, err := payload.File.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()
	url = fmt.Sprintf("%s/object/%s/%s", url, payload.Folder, this.Util.GenerateFileName())
	req, err := http.NewRequest(http.MethodPost, url, file)

	if err != nil {
		return "", err
	}

	token := fmt.Sprintf("Bearer %s", this.Config.ServiceRoleKey)
	mimeType := fmt.Sprintf("%s/*", payload.Type)
	req.Header.Set("apikey", this.Config.PublishableKey)
	req.Header.Set(fiber.HeaderContentType, mimeType)
	req.Header.Set(fiber.HeaderAuthorization, token)

	client := http.Client{}
	res, err := client.Do(req)

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	stringBody := string(body)

	return stringBody, nil
}
