package v1

import (
	"marketplace-backend/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AddTagsToListing adds one or more tags to a listing, ensuring validity and uniqueness.
func AddTagsToListing(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			ListingID uint     `json:"listingId"`
			Tags      []string `json:"tags"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
			return
		}

		// Verify that the listing exists
		var listing model.Listing
		if err := db.First(&listing, request.ListingID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Listing not found"})
			return
		}

		// Process each tag: create new tags as needed and associate them with the listing
		for _, tagName := range request.Tags {
    	var tag model.Tag
   		 if err := db.Where(model.Tag{Name: tagName}).FirstOrCreate(&tag).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create or find tag"})
        return
    }
    // Associate tag with listing if not already associated
    if err := db.Model(&listing).Association("Tags").Append(&tag); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to associate tag with listing"})
        return
    }
}


// GetListingsByTag retrieves listings associated with a specific tag, optimizing query efficiency.
func GetListingsByTag(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tagName := c.Param("tagName")

		// Efficiently retrieve listings by tag name using JOIN
		var listings []model.Listing
		if err := db.Joins("JOIN listing_tags ON listing_tags.listing_id = listings.id").
			Joins("JOIN tags ON tags.id = listing_tags.tag_id").
			Where("tags.name = ?", tagName).
			Find(&listings).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve listings"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"listings": listings})
	}
}
