package model

import (
	"gorm.io/gorm"
	"time"
)

// Listing represents an item or service offered on the marketplace.
type Listing struct {
	gorm.Model
	UserID      uint      `gorm:"not null"`
	Title       string    `gorm:"not null"`
	Description string    `gorm:"type:text"`
	Price       float64   `gorm:"not null"`
	Category    string    `gorm:"not null"`
	Tags 		[]Tag 	  `gorm:"many2many:listing_tags;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
