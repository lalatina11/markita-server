package routes

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/lalatina11/markita.git/src/lib/response"
)

type MediaForm struct {
	media byte
}

func UploadRoutes(api fiber.Router) *fiber.Router {
	r := api.Group("/upload")

	r.Get("/", func(ctx fiber.Ctx) error {
		return response.ErrorResponse(ctx, nil, nil)
	})

	r.Post("/", func(c fiber.Ctx) error {
		// form := new(MediaForm)
		// if err := c.Bind().Body(form); err != nil {
		// 	msg := err.Error()
		// 	return response.ErrorResponse(c, &msg, nil)
		// }
		// fmt.Println("form retrieved\n")
		// fmt.Println(form)
		media, err := c.FormFile("media")
		if err != nil {
			msg := err.Error()
			return response.ErrorResponse(c, &msg, nil)
		}
		isImage := strings.Contains(media.Header.Get("Content-Type"), "image")
		isVideo := strings.Contains(media.Header.Get("Content-Type"), "video")
		if !isImage && !isVideo {
			msg := "Please insert a video or image"
			return response.ErrorResponse(c, &msg, nil)
		}
		mediaType := "image"
		if isVideo {
			mediaType = "video"
		}
		fmt.Println(mediaType)
		fmt.Println(media.Header)
		fmt.Println(media.Filename)
		fmt.Println(media.Size)
		return response.SuccessResponse[any](c, nil, nil, nil)
	})

	return &r
}
