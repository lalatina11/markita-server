package service

import (
	"encoding/json"

	"github.com/lalatina11/markita.git/src/error/service_error"
	"github.com/lalatina11/markita.git/src/lib/payload"
	supabaseresponse "github.com/lalatina11/markita.git/src/lib/response/supabase_response"
)

type UploadService struct {
	SupabaseService *SupabaseService
}

func NewUploadService() *UploadService {
	SupabaseService := NewSupabaseService()
	return &UploadService{SupabaseService}
}

func (this *UploadService) UploadFile(payload *payload.FileUploadPayload) (*supabaseresponse.StorageUploadSuccessResponse, *service_error.ServiceError) {
	stringRes, err := this.SupabaseService.StorageUploadFile(payload)

	if err != nil {
		return nil, service_error.Create(500, err.Error())
	}

	successResponse := new(supabaseresponse.StorageUploadSuccessResponse)

	if err := json.Unmarshal([]byte(stringRes), successResponse); err == nil && successResponse.IsSuccess() {
		return successResponse, nil
	}

	errorResponse := new(supabaseresponse.StorageUploadErrorResponse)

	if _err := json.Unmarshal([]byte(stringRes), errorResponse); _err == nil && errorResponse.IsError() {
		return nil, errorResponse.ToServiceError()
	}

	return nil, service_error.InternalServerError()

}
