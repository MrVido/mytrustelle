package model

import (
	"gorm.io/gorm"
)

// Tag represents a keyword or label that categorizes a listing.
// Enhancements include usage statistics, hierarchy, and search optimization.
type Tag struct {
	gorm.Model
	Name        string    `gorm:"uniqueIndex;not null"` // Unique index to avoid duplicate tags
	Description string    `gorm:"type:text"`            // Description of what the tag represents
	Listings    []Listing `gorm:"many2many:listing_tags;"` // The listings associated with this tag
	UseCount    int       `gorm:"default:0"`            // How often the tag is used across listings
	ParentID    *uint     `gorm:"index"`                // Allows hierarchical structuring of tags (optional)
	Children    []Tag     `gorm:"foreignKey:ParentID"`  // Child tags for hierarchical organization
	Synonyms    []Tag     `gorm:"many2many:tag_synonyms;"` // Other tags that are considered synonyms
}

// ListingTag represents the many-to-many relationship between Listings and Tags.
// Including timestamps can help track when tags are added to listings.
type ListingTag struct {
	ListingID uint `gorm:"primaryKey"`
	TagID     uint `gorm:"primaryKey"`
	CreatedAt time.Time // When the tag was associated with the listing
	UpdatedAt time.Time // Last time the association was updated
}

// TagSynonym represents the many-to-many relationship between Tags themselves to handle synonyms.
type TagSynonym struct {
	TagID      uint `gorm:"primaryKey"`
	SynonymTagID uint `gorm:"primaryKey"`
}

// Key Enhancements:

// 1. **Tag Descriptions**: Providing a `Description` for tags can help users understand the context and proper use of a tag.

// 2. **Usage Statistics**: `UseCount` tracks the popularity or relevance of a tag, which can be useful for tag cleanup, recommendations, or trending tag displays.

// 3. **Hierarchical Tags**: `ParentID` and `Children` allow for nested tags, enabling more structured and navigable categorization.

// 4. **Synonyms**: The `Synonyms` relation allows for linking tags that mean the same thing, enhancing search and categorization without needing to unify tag names immediately.

// 5. **Many-to-Many Enhancements**: The `ListingTag` structure with `CreatedAt` and `UpdatedAt` timestamps provides insights into tagging trends over time.

// These enhancements aim to make the tagging system more versatile and informative, supporting a wide range of use cases from simple categorization to complex taxonomies and search optimizations within your marketplace platform.
