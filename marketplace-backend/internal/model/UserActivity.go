package model

import (
	"gorm.io/gorm"
	"time"
)

// UserActivity logs activities performed by users for analytics.
type UserActivity struct {
	gorm.Model
	UserID    uint      `gorm:"not null;index"` // The user performing the activity
	Activity  string    `gorm:"not null"`       // Description of the activity
	Timestamp time.Time `gorm:"default:CURRENT_TIMESTAMP"` // When the activity occurred
}

// Enhancements:
// - Insights into user behavior.
// - Basis for personalized recommendations and services.
