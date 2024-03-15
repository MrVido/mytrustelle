package v1

import (
    "marketplace-backend/internal/model"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// GetUserRewards retrieves the list of rewards earned by a user, with checks for expiration.
func GetUserRewards(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        userID, exists := c.Get("userID")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            return
        }

        // Include a check to filter out expired rewards if necessary
        var rewards []model.UserReward
        if result := db.Where("user_id = ? AND (expires_at > ? OR expires_at IS NULL)", userID, time.Now()).Find(&rewards); result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
            return
        }

        c.JSON(http.StatusOK, rewards)
    }
}

// RedeemReward processes a user's request to redeem a reward.
func RedeemReward(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        userID, exists := c.Get("userID")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            return
        }

        var redeemRequest struct {
            RewardID uint `json:"rewardId"`
        }
        if err := c.ShouldBindJSON(&redeemRequest); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
            return
        }

        var reward model.UserReward
        if err := db.Where("id = ? AND user_id = ?", redeemRequest.RewardID, userID).First(&reward).Error; err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Reward not found or not eligible for redemption"})
            return
        }

        // Ensure the reward is not expired and not already redeemed
        if reward.Redeemed || (reward.ExpiresAt != nil && reward.ExpiresAt.Before(time.Now())) {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Reward has already been redeemed or is expired"})
            return
        }

        // Mark the reward as redeemed
        reward.Redeemed = true
        reward.RedeemedAt = time.Now()
        if result := db.Save(&reward); result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not redeem reward"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Reward redeemed successfully"})
    }
}
