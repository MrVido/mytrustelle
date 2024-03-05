package v1

import (
	"marketplace-backend/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// PostReview allows a user to leave a review for another user.
func PostReview(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var review model.Review
		if err := c.ShouldBindJSON(&review); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Add additional validation if necessary (e.g., check if the reviewer and reviewee are different)

		if result := db.Create(&review); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusCreated, review)
	}
}

// GetReviews fetches reviews for a given user.
func GetReviews(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reviews []model.Review
		userID := c.Param("userID") // Assumes userID is passed as a URL parameter

		if result := db.Where("reviewee_id = ?", userID).Find(&reviews); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, reviews)
	}
}
