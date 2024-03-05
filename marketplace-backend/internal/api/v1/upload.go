package v1

import (
	"marketplace-backend/internal/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UploadFile handles the file upload process, including validation and storage.
func UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
		return
	}
	defer file.Close()

	if err := util.ValidateFileUpload(header); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// After validation, upload the file to your storage solution (e.g., Amazon S3)
	// This step depends on the SDK or API of the storage service you're using.
	// Example: util.UploadToS3(file)

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}
