package model

import (
	"gorm.io/gorm"
)

// APIKey manages API keys issued to users or third-party integrations.
type APIKey struct {
	gorm.Model
	UserID uint   `gorm:"not null;index"` // The user to whom the API key is issued
	Key    string `gorm:"not null;uniqueIndex"` // The API key itself
	Active bool   `gorm:"default:true"`         // Whether the API key is active
}

// Enhancements:
// - Secure integration with third-party services.
// - Control access with the ability to deactivate keys.
