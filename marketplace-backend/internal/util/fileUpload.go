package util

import (
	"errors"
	"mime/multipart"
	"net/http"
)

const maxUploadSize = 10 << 20 // 10 MB

// ValidateFileUpload checks the MIME type and size of the uploaded file.
func ValidateFileUpload(header *multipart.FileHeader) error {
	if header.Size > maxUploadSize {
		return errors.New("file size exceeds the allowable limit")
	}

	fileType := header.Header.Get("Content-Type")
	switch fileType {
	case "image/jpeg", "image/png":
		return nil
	default:
		return errors.New("invalid file type")
	}
}
