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
// Enhanced GetUserRecommendations with caching and user preferences consideration.
func GetUserRecommendations(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        userID, exists := c.Get("userID")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
            return
        }

        uid, ok := userID.(uint)
        if !ok {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error processing userID"})
            return
        }

        // Example: Check for cached recommendations first
        var recommendations []model.Listing
        cacheKey := fmt.Sprintf("user:%d:recommendations", uid)
        if err := util.GetCache(cacheKey, &recommendations); err == nil && len(recommendations) > 0 {
            c.JSON(http.StatusOK, gin.H{"recommendations": recommendations})
            return
        }

        // Fetch user preferences and generate recommendations based on them
        preferences, err := util.GetUserPreferences(db, uid)
        if err != nil {
            log.Printf("Error fetching user preferences for userID %d: %v", uid, err)
            // Consider whether to fail or proceed with default recommendations in case of error
        }

        recommendations, err = util.GenerateRecommendations(db, uid, preferences)
        if err != nil {
            log.Printf("Error generating recommendations for userID %d: %v", uid, err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate recommendations"})
            return
        }

        // Cache the recommendations to avoid regenerating them on every request
        util.SetCache(cacheKey, recommendations, 30*time.Minute)

        c.JSON(http.StatusOK, gin.H{"recommendations": recommendations})
    }
}
