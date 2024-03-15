package model

import (
	"gorm.io/gorm"
	"time"
)

// UserReward represents the association between a user and earned rewards in the loyalty program.
type UserReward struct {
	gorm.Model
	UserID    uint          // ID of the user who earned the reward
	RewardID  uint          // ID of the reward
	Reward    LoyaltyReward // The associated reward, using Gorm's foreign key relation
	Redeemed  bool          // Indicates if the reward has been redeemed
	RedeemedAt *time.Time   // Timestamp of when the reward was redeemed
	ExpiresAt  *time.Time   // Expiration date of the reward; nil means no expiration
}

// LoyaltyReward represents a reward that can be earned by users.
type LoyaltyReward struct {
	gorm.Model
	Name        string     // Name of the reward
	Description string     // Description of the reward
	Points      int        // Points cost of the reward
	ExpiresAt   *time.Time // Expiration date; nil means no expiration
}

// TableName overrides the table name used by UserReward to prevent Gorm from using the pluralized version.
func (UserReward) TableName() string {
	return "user_reward"
}
