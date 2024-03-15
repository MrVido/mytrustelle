package model

import (
	"gorm.io/gorm"
	"time"
)

// PasswordResetToken manages password recovery processes.
type PasswordResetToken struct {
	gorm.Model
	UserID    uint      `gorm:"not null;index"` // ID of the user requesting password reset
	Token     string    `gorm:"not null;uniqueIndex"` // Reset token
	ExpiresAt time.Time `gorm:"not null"`             // Expiration time of the token
}

// Enhancements:
// - Secure management of password reset requests.
// - Automatic expiration of tokens to mitigate risk.
