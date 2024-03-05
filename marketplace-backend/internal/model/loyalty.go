package model

import (
	"gorm.io/gorm"
	"time"
)

// LoyaltyReward represents a reward that users can earn through the loyalty program.
type LoyaltyReward struct {
	gorm.Model
	Name        string    `gorm:"not null"` // Name of the reward
	Description string    `gorm:"type:text;not null"` // Description of the reward
	Points      int       `gorm:"not null"` // Points required to redeem the reward
	ExpiresAt   time.Time // Optional expiration date for the reward
}

// UserReward represents the association between a user and earned rewards.
type UserReward struct {
	gorm.Model
	UserID      uint      `gorm:"not null"`
	RewardID    uint      `gorm:"not null"`
	Reward      LoyaltyReward `gorm:"foreignKey:RewardID"`
	Redeemed    bool      `gorm:"not null;default:false"` // Whether the reward has been redeemed
	RedeemedAt  time.Time // When the reward was redeemed
}
