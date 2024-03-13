package v1

import (
    "marketplace-backend/internal/model"
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// GetUserRewards lists the rewards earned by a user.
func GetUserRewards(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        userID, exists := c.Get("userID")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            return
        }

        var rewards []model.UserReward
        if result := db.Where("user_id = ?", userID).Find(&rewards); result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
            return
        }

        c.JSON(http.StatusOK, rewards)
    }
}

// RedeemReward allows a user to redeem a reward.
func RedeemReward(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var redeemRequest struct {
            RewardID uint `json:"rewardId"`
        }
        userID, exists := c.Get("userID")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            return
        }

        if err := c.ShouldBindJSON(&redeemRequest); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
            return
        }

        // Validate reward exists and belongs to the user
        var reward model.UserReward
        if err := db.Where("id = ? AND user_id = ?", redeemRequest.RewardID, userID).First(&reward).Error; err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Reward not found or does not belong to user"})
            return
        }

        // Further checks can be added here to ensure the reward has not been redeemed, is not expired, etc.

        // Update the reward as redeemed
        reward.Redeemed = true
        if result := db.Save(&reward); result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not redeem reward"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Reward redeemed successfully"})
    }
}
