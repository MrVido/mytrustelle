package model

import (
	"gorm.io/gorm"
	"time"
)

// Review represents feedback left by a user for another user.
type Review struct {
	gorm.Model
	ReviewerID uint      `gorm:"not null"` // ID of the user leaving the review
	RevieweeID uint      `gorm:"not null"` // ID of the user being reviewed
	ListingID  uint      `gorm:"not null"` // Associated listing
	Rating     int       `gorm:"not null"` // Rating given by the reviewer
	Comment    string    // Optional comment
	CreatedAt  time.Time
}
