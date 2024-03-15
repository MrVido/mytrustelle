package model

import (
	"gorm.io/gorm"
)

// Category represents a category or subcategory within the marketplace.
// Enhancements include support for hierarchical categories and SEO metadata.
type Category struct {
	gorm.Model
	Name        string   `gorm:"not null;uniqueIndex"` // Unique name of the category
	Description string   `gorm:"type:text"`            // Description of the category
	ParentID    *uint    `gorm:"index"`                // Allows for nesting categories. Null for top-level categories.
	Children    []Category `gorm:"foreignKey:ParentID"`// Nested subcategories
	Slug        string   `gorm:"not null;uniqueIndex"` // URL-friendly version of the name for SEO
	MetaTitle   string   // Optional SEO meta title
	MetaDescription string `gorm:"type:text"`          // Optional SEO meta description for better search engine visibility
}

// Key Enhancements:

// 1. **Hierarchical Categories**: The `ParentID` field allows categories to have subcategories, creating a tree-like structure. This is useful for organizing listings into broader categories and more specific subcategories.

// 2. **SEO Optimization**: `Slug`, `MetaTitle`, and `MetaDescription` fields are added to enhance SEO. The `Slug` is a URL-friendly version of the category name, which can be used in the routing to make URLs more descriptive and SEO-friendly. `MetaTitle` and `MetaDescription` provide additional metadata that can be used in HTML headers to improve the visibility of category pages in search engines.

// 3. **Self-referencing Children**: The `Children` slice of `Category` struct allows for easy retrieval of subcategories directly associated with a category, facilitating nested category display and navigation.

// This enhanced `Category` model provides a flexible and scalable way to organize listings in your marketplace, improving navigation and search engine rankings.
