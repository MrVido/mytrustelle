package model

import (
	"gorm.io/gorm"
)

// UserSetting stores user-specific preferences and settings.
type UserSetting struct {
	gorm.Model
	UserID      uint   `gorm:"not null;uniqueIndex"` // Reference to the User model
	Preferences string `gorm:"type:text"`            // JSON-encoded preferences (e.g., notification settings)
}

// Enhancements:
// - Personalization of user experience.
// - Flexible storage of varied preference types.
