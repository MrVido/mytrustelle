package model

import (
	"gorm.io/gorm"
)

// ContactMessage stores messages from contact forms for support and inquiries.
type ContactMessage struct {
	gorm.Model
	Name    string `gorm:"not null"` // Name of the person contacting
	Email   string `gorm:"not null"` // Contact email address
	Message string `gorm:"type:text;not null"` // Message content
}

// Enhancements:
// - Streamlined support ticket creation.
// - Automated responses or routing based on message content.
