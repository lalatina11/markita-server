package handler

import (
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/lib/payload"
	"github.com/lalatina11/markita.git/src/lib/response"
	"github.com/lalatina11/markita.git/src/service"
)

type UploadHandler struct {
	UploadService *service.UploadService
}

func NewUploadHandler() *UploadHandler {
	UploadService := service.NewUploadService()
	return &UploadHandler{UploadService}
}

func (this *UploadHandler) Upload(c fiber.Ctx) error {
	folder := c.FormValue("folder", "")

	media, err := c.FormFile("media")
	if err != nil {
		msg := "Please insert an Image or Video"
		return response.ErrorResponse(c, &msg, nil)
	}
	mimeType := media.Header.Get("Content-Type")
	isImage := strings.Contains(mimeType, "image")
	isVideo := strings.Contains(mimeType, "video")
	if !isImage && !isVideo {
		msg := "Please insert a video or image"
		return response.ErrorResponse(c, &msg, nil)
	}
	mediaType := "image"
	if isVideo {
		mediaType = "video"
	}
	payload := payload.FileUploadPayload{
		File:     media,
		Type:     mediaType,
		Name:     media.Filename,
		Folder:   folder,
		MimeType: mimeType,
	}

	res, serviceError := this.UploadService.UploadFile(&payload)

	if serviceError != nil {
		return serviceError.ToResponse(c)
	}

	return response.SuccessResponse(c, nil, res.ToJson(mediaType), nil)
}
