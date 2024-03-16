package util

import (
	"errors"
	"fmt"
	"mime/multipart"
)

const (
	maxUploadSize = 10 << 20 // 10 MB, adjust as needed
)

// Update here to include all your allowed MIME types
var allowedMimeTypes = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
	"image/webp": true,
	"image/heic": true,
}

// ValidateFileUpload checks the MIME type and size of the uploaded file with detailed errors.
func ValidateFileUpload(header *multipart.FileHeader) error {
	if header.Size > maxUploadSize {
		return fmt.Errorf("file size (%d bytes) exceeds the allowable limit (%d bytes)", header.Size, maxUploadSize)
	}

	fileType := header.Header.Get("Content-Type")
	if _, ok := allowedMimeTypes[fileType]; !ok {
		// Use the helper function to generate a formatted string of allowed MIME types
		return fmt.Errorf("file type '%s' is not supported. Allowed types are: %s", fileType, getFormattedAllowedTypes())
	}

	return nil
}

// getFormattedAllowedTypes returns a formatted string of allowed MIME types for error messages.
func getFormattedAllowedTypes() string {
	types := ""
	for mimeType := range allowedMimeTypes {
		types += mimeType + ", "
	}
	// Trim the trailing comma and space for a cleaner message
	types = types[:len(types)-2]
	return types
}
