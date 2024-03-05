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
		userID, _ := c.Get("userID") // Assumes userID is extracted from JWT token

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
		if err := c.ShouldBindJSON(&redeemRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Logic to mark the reward as redeemed for the user
		// Ensure the user has enough points and the reward hasn't expired or already been redeemed.

		c.JSON(http.StatusOK, gin.H{"message": "Reward redeemed successfully"})
	}
}
