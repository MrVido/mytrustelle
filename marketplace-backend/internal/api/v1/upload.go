package v1

import (
    "marketplace-backend/internal/util"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "path/filepath"
)

// UploadFile handles secure and validated file uploads.
func UploadFile(c *gin.Context) {
    file, header, err := c.Request.FormFile("upload")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to receive file"})
        return
    }
    defer file.Close()

    if err := util.ValidateFileUpload(header); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    uniqueFileName := uuid.New().String() + filepath.Ext(header.Filename)
    uploadedFileURL, err := util.UploadToStorage(file, uniqueFileName)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "fileUrl": uploadedFileURL})
}
