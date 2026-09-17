package payload

import "mime/multipart"

type FileUploadPayload struct {
	File     *multipart.FileHeader
	Type     string
	Name     string
	Folder   string
	MimeType string
}
