package v1

import (
	"net/http"
	"marketplace-backend/internal/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetUserRecommendations retrieves personalized listing recommendations for a user.
func GetUserRecommendations(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("userID") // Assumes userID is extracted from JWT token

		recommendations, err := util.GenerateRecommendations(db, userID.(uint))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to generate recommendations"})
			return
		}

		c.JSON(http.StatusOK, recommendations)
	}
}
