package model

import (
	"gorm.io/gorm"
	"time"
)

// Image represents an image associated with a listing.
type Image struct {
	gorm.Model
	ListingID uint      `gorm:"not null"` // Foreign key to the Listing
	URL       string    `gorm:"not null"` // URL to access the image
	CreatedAt time.Time
	UpdatedAt time.Time
}
