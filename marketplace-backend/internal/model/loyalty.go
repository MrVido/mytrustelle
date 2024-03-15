package model

import (
	"gorm.io/gorm"
)

// LoyaltyReward represents a reward that users can earn through the loyalty program.
type LoyaltyReward struct {
	gorm.Model
	Name           string `gorm:"not null"` // Name of the reward
	Description    string `gorm:"type:text;not null"` // Description of the reward
	PointsRequired int    `gorm:"not null"` // Points required to redeem the reward
	ExpiresAt      *time.Time // Optional expiration date for the reward
	Availability   int    `gorm:"default:-1"` // Number of rewards available (-1 for unlimited)
	Tier           string `gorm:"type:varchar(20)"` // Tier level required to redeem the reward
}

// UserReward represents the association between a user and earned rewards.
type UserReward struct {
	gorm.Model
	UserID     uint         `gorm:"not null;index"` // Indexed for faster lookups by UserID
	RewardID   uint         `gorm:"not null"`
	Reward     LoyaltyReward `gorm:"foreignKey:RewardID"`
	Redeemed   bool         `gorm:"not null;default:false"` // Whether the reward has been redeemed
	RedeemedAt *time.Time   // When the reward was redeemed, if applicable
}

// UserLoyaltyProfile represents a user's status and points within the loyalty program.
type UserLoyaltyProfile struct {
	gorm.Model
	UserID    uint    `gorm:"uniqueIndex;not null"`
	Points    int     `gorm:"default:0"` // Total loyalty points earned by the user
	Tier      string  `gorm:"type:varchar(20);default:'bronze'"` // Current tier level of the user
	LastAward *time.Time // Last time points were awarded, for periodic reward calculations
}

// Key Enhancements:

// 1. **Tier Levels**: `Tier` in `LoyaltyReward` allows for tier-based rewards, encouraging users to achieve higher levels within the loyalty program.

// 2. **Limited Availability**: `Availability` in `LoyaltyReward` tracks how many of such rewards are available, adding exclusivity and urgency to redeem.

// 3. **User Loyalty Profile**: Introducing `UserLoyaltyProfile` to manage a user's loyalty points, tier level, and track the last award date for periodic rewards.

// 4. **Redeem Tracking**: `Redeemed` and `RedeemedAt` in `UserReward` provide detailed tracking of reward redemption status and timing.

// 5. **Database Indexes**: Adding an index on `UserID` in `UserReward` improves the efficiency of querying a user's rewards.

// These enhancements aim to create a more engaging and structured loyalty program, offering users clear paths for earning rewards and incentivizing continued participation in the marketplace.