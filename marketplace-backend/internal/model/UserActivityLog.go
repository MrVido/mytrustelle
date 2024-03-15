package model

import (
	"gorm.io/gorm"
	"time"
)

// UserActivityLog details logs of user actions for security and analytics.
type UserActivityLog struct {
	gorm.Model
	UserID    uint      `gorm:"not null;index"` // User performing the action
	Action    string    `gorm:"not null"`       // Description of the action
	Timestamp time.Time `gorm:"default:CURRENT_TIMESTAMP"` // Time of the action
}

// Enhancements:
// - Insights into user engagement and potential security breaches.
// - Basis for behavioral analytics and personalized experiences.
