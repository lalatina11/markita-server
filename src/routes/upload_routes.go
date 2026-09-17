package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/handler"
)

type MediaForm struct {
	media byte
}

func UploadRoutes(api fiber.Router) *fiber.Router {
	r := api.Group("/upload")

	handler := handler.NewUploadHandler()

	r.Post("/", handler.Upload)

	return &r
}
