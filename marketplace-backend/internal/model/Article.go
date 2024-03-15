package model

import (
	"gorm.io/gorm"
	"time"
)

// Article represents dynamic content such as help articles, FAQs, or blogs.
type Article struct {
	gorm.Model
	Title      string    `gorm:"not null;uniqueIndex"` // Article title
	Content    string    `gorm:"type:text;not null"`   // Article content (HTML or Markdown)
	AuthorID   uint      `gorm:"not null"`             // ID of the author (reference to a User model)
	Published  bool      `gorm:"default:false"`        // Whether the article is published
	PublishAt  time.Time // Scheduled publish time
	Tags       []Tag     `gorm:"many2many:article_tags;"` // Tags for categorizing articles
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// Enhancements:
// - SEO optimization through tags and structured content.
// - Scheduled publishing for content planning.
