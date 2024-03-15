package model

import (
	"gorm.io/gorm"
	"time"
)

// UserSession tracks active user sessions for security and analytics.
type UserSession struct {
	gorm.Model
	UserID    uint      `gorm:"not null;index"` // ID of the logged-in user
	SessionID string    `gorm:"not null;uniqueIndex"` // Session identifier
	ExpiresAt time.Time `gorm:"not null"`             // Expiration time of the session
}

// Enhancements:
// - Auto-expiration of stale sessions.
// - Support for session invalidation on logout or password change.
