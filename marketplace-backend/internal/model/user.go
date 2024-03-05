package model

import (
	"gorm.io/gorm"
	"time"
)

// User represents a user in the marketplace.
type User struct {
	gorm.Model
	Username     string `gorm:"uniqueIndex;not null"`
	Email        string `gorm:"uniqueIndex;not null"`
	PhoneNumber  string `gorm:"not null"`
	PasswordHash string `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Role		 string `gorm:"default:user"` // "user" or "admin"
	EmailPreferences struct {
        Marketing bool `gorm:"default:true"` // Opt-in to marketing emails by default
        Updates   bool `gorm:"default:true"` // Opt-in to platform updates emails by default
    } `gorm:"embedded"`
}
