package model

import (
	"gorm.io/gorm"
)

// ContentPage stores editable content pages (e.g., "About Us", "Terms of Service").
type ContentPage struct {
	gorm.Model
	Title   string `gorm:"not null;uniqueIndex"` // Page title
	Slug    string `gorm:"not null;uniqueIndex"` // URL-friendly identifier
	Content string `gorm:"type:text;not null"`   // Page content (HTML or Markdown)
}

// Enhancements:
// - Easy content management without code deployment.
// - SEO-friendly URLs through slugs.
