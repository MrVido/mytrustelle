package model

import (
	"gorm.io/gorm"
)

// Tag represents a keyword or label that categorizes a listing.
type Tag struct {
	gorm.Model
	Name     string    `gorm:"uniqueIndex;not null"`
	Listings []Listing `gorm:"many2many:listing_tags;"`
}

// ListingTag represents the many-to-many relationship between Listings and Tags.
type ListingTag struct {
	ListingID uint `gorm:"primaryKey"`
	TagID     uint `gorm:"primaryKey"`
}
