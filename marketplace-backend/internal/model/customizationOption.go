package model

import (
	"gorm.io/gorm"
)

// CustomizationOption allows sellers to offer customizations for their products.
type CustomizationOption struct {
	gorm.Model
	ListingID uint   `gorm:"not null"` // Associated listing
	Name      string `gorm:"not null"` // Name of the customization (e.g., color, size)
	Choices   string `gorm:"type:text;not null"` // JSON-encoded choices available for this customization
}

// Enhancements:
// - Flexibility for sellers to define product variants.
// - Enhanced buyer experience with personalized products.
