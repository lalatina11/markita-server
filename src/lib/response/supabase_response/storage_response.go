package supabaseresponse

import (
	"strconv"

	"github.com/lalatina11/markita.git/src/error/service_error"
)

type StorageUploadSuccessResponse struct {
	Key string // url
	Id  string // uuid
}

type StorageUploadSuccessJsonResponse struct {
	Url  string `json:"url"`
	Type string `json:"type"`
}

func (this *StorageUploadSuccessResponse) IsSuccess() bool {
	return this.Key != ""
}
func (this *StorageUploadSuccessResponse) ToJson(mediaType string) *StorageUploadSuccessJsonResponse {
	return &StorageUploadSuccessJsonResponse{
		Url:  this.Key,
		Type: mediaType,
	}
}

type StorageUploadErrorResponse struct {
	StatusCode string `json:"statusCode"`
	Message    string `json:"message"`
}

func (this *StorageUploadErrorResponse) ToServiceError() *service_error.ServiceError {
	code, err := strconv.Atoi(this.StatusCode)
	if err != nil {
		code = 400
	}
	return service_error.CreateServiceError(code, this.Message)
}
func (this *StorageUploadErrorResponse) IsError() bool {
	return this.Message != ""
}
