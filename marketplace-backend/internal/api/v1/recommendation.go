package v1

import (
    "net/http"
    "marketplace-backend/internal/model" // Assume this package contains your models
    "marketplace-backend/internal/util"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "log" // Added for logging purposes
)

// GetUserRecommendations retrieves personalized listing recommendations for a user.
func GetUserRecommendations(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Validate and extract the userID from the context, ensuring authorization
        userID, exists := c.Get("userID")
        if !exists || userID == nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized or ID missing"})
            return
        }

        // Cast userID to the expected type (uint) and validate
        uid, ok := userID.(uint)
        if !ok {
            log.Println("Error casting userID to uint, incorrect type provided")
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
            return
        }

        // Call the utility function to generate recommendations. 
        // This function is assumed to be complex and potentially leveraging machine learning.
        recommendations, err := util.GenerateRecommendations(db, uid)
        if err != nil {
            log.Printf("Error generating recommendations for userID %d: %v", uid, err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to generate recommendations"})
            return
        }

        // Optional: Further process or filter recommendations before sending to the client
        // For example, exclude already seen items or apply additional user-specific filters

        c.JSON(http.StatusOK, gin.H{"recommendations": recommendations})
    }
}
