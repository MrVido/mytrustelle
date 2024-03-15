package v1

import (
    "marketplace-backend/internal/util"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "path/filepath"
)

// UploadFile handles the file upload process, adding security, validation, and storage.
func UploadFile(c *gin.Context) {
    file, header, err := c.Request.FormFile("upload")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to receive file"})
        return
    }
    defer file.Close()

    // Validate the file type and size
    if err := util.ValidateFileUpload(header); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Generate a unique file name to prevent overwrites
    uniqueFileName := uuid.New().String() + filepath.Ext(header.Filename)

    // After validation, upload the file to your storage solution (e.g., Amazon S3, Google Cloud Storage)
    // Assuming util.UploadToStorage is a function that abstracts the details of the storage service
    uploadedFileURL, err := util.UploadToStorage(file, uniqueFileName)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
        return
    }

    // Return the URL or identifier of the uploaded file to the client
    c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "fileUrl": uploadedFileURL})
}
