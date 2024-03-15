package model

import (
	"gorm.io/gorm"
)

// InventoryItem tracks stock levels for listings.
type InventoryItem struct {
	gorm.Model
	ListingID  uint `gorm:"not null;uniqueIndex"` // Associated listing
	StockLevel int  `gorm:"not null"`             // Quantity available
}

// Enhancements:
// - Real-time stock updates to prevent overselling.
// - Notifications for low stock levels.
