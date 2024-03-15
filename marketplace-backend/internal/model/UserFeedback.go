package model

import (
	"gorm.io/gorm"
)

// UserFeedback captures user feedback for improving user experience and support.
type UserFeedback struct {
	gorm.Model
	UserID  uint   `gorm:"not null;index"` // ID of the user providing feedback
	Content string `gorm:"type:text;not null"` // Feedback content
}

// Enhancements:
// - Analysis of feedback trends to guide product improvements.
// - Support for categorizing feedback types for targeted responses.
