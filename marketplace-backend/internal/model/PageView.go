package model

import (
	"gorm.io/gorm"
	"time"
)

// PageView tracks views for analytics, helping understand content popularity and user engagement.
type PageView struct {
	gorm.Model
	PageURL   string    `gorm:"not null"` // URL of the viewed page
	VisitorID string    `gorm:"not null"` // Identifier for the visitor (could be session ID or user ID if logged in)
	ViewedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"` // When the page was viewed
}

// Enhancements:
// - Detailed analysis of page popularity and visitor engagement.
// - Support for A/B testing by tracking variant interactions.
