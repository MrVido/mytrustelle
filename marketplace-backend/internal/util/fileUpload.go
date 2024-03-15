package util

import (
	"fmt"
	"mime/multipart"
)

// Configurable constants for easier adjustments
const (
	maxUploadSize = 10 << 20 // 10 MB, adjust as needed
)

var allowedMimeTypes = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
	// Add more MIME types as needed
}

// ValidateFileUpload checks the MIME type and size of the uploaded file with detailed errors.
func ValidateFileUpload(header *multipart.FileHeader) error {
	if header.Size > maxUploadSize {
		return fmt.Errorf("file size (%d bytes) exceeds the allowable limit (%d bytes)", header.Size, maxUploadSize)
	}

	fileType := header.Header.Get("Content-Type")
	if !allowedMimeTypes[fileType] {
		allowedTypes := getFormattedAllowedTypes()
		return fmt.Errorf("invalid file type '%s'. Allowed types are: %s", fileType, allowedTypes)
	}

	return nil
}

// getFormattedAllowedTypes returns a formatted string of allowed MIME types for error messages.
func getFormattedAllowedTypes() string {
	types := ""
	for mimeType := range allowedMimeTypes {
		types += mimeType + ", "
	}
	// Remove the trailing comma and space
	return types[:len(types)-2]
}
