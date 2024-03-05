package v1

import (
	"fmt"
	"marketplace-backend/internal/model"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UploadImage handles the uploading of images for a listing.
func UploadImage(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Assuming listingID is passed as a form value
		listingID := c.PostForm("listingID")

		// Here you should handle the file upload process. The specific implementation will vary
		// depending on whether you're storing images locally or using a cloud service.
		// This is a placeholder for the file upload logic.

		// Assume `uploadedURL` is the URL to the uploaded image
		uploadedURL := "http://example.com/path/to/image.jpg"

		// Create and save the image record in the database
		image := model.Image{ListingID: uint(listingID), URL: uploadedURL}
		if result := db.Create(&image); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Image uploaded successfully", "image": image})
	}
}
