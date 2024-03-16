package v1

import (
	"marketplace-backend/internal/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// PostReview allows a user to leave a review for another user with added validations.
func PostReview(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var review model.Review
		if err := c.ShouldBindJSON(&review); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review format"})
			return
		}

		// Authentication: Ensure the reviewer is logged in
		reviewerID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Validation: Ensure the reviewer and reviewee are different
		if review.ReviewerID == review.RevieweeID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Reviewer and reviewee cannot be the same"})
			return
		}

		// Additional validation: Check if the reviewee exists
		var user model.User
		if err := db.First(&user, review.RevieweeID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Reviewee does not exist"})
			return
		}

		if review.Rating < 1 || review.Rating > 5 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Rating must be between 1 and 5"})
			return
		}

		// Set the reviewer ID to the current user's ID
		review.ReviewerID = reviewerID.(uint)

		if result := db.Create(&review); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to post review"})
			return
		}

		// After successfully creating the review, update aggregate rating
		if err := updateAggregateRating(db, review.RevieweeID); err != nil {
			log.Printf("Failed to update aggregate rating: %v", err)
			// Consider how you want to handle this error. For simplicity, we're logging it but not interrupting the response.
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Review posted successfully", "review": review})
	}
}

// GetReviews fetches reviews for a given user with validation.
func GetReviews(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDParam := c.Param("userID")
		userID, err := strconv.ParseUint(userIDParam, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		var reviews []model.Review
		if result := db.Where("reviewee_id = ?", userID).Find(&reviews); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reviews"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"reviews": reviews})
	}
}
func updateAggregateRating(db *gorm.DB, revieweeID uint) error {
	var averageRating float64
	db.Model(&model.Review{}).Where("reviewee_id = ?", revieweeID).Select("AVG(rating)").Row().Scan(&averageRating)

	// Assume a User model with an AggregateRating field, if not, adjust as necessary.
	return db.Model(&model.User{}).Where("id = ?", revieweeID).Update("aggregate_rating", averageRating).Error
}

