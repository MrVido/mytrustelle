package model

import (
	"gorm.io/gorm"
)

// Image represents an image associated with a listing.
// It includes soft deletes, additional metadata for accessibility and SEO, and flags for primary image identification.
type Image struct {
	gorm.Model
	ListingID uint           `gorm:"index;not null"` // Ensures faster queries by listing ID
	URL       string         `gorm:"not null"`       // The URL where the image is stored
	AltText   string         `gorm:"type:text"`      // Alternative text for the image, crucial for SEO and accessibility
	Caption   string         `gorm:"type:text"`      // A caption for the image, providing more context
	IsPrimary bool           `gorm:"default:false"`  // Indicates if this is the primary image for the listing
	DeletedAt gorm.DeletedAt `gorm:"index"`          // Enables soft delete functionality
}

// Enhancements include:
// 1. **Image Processing & Optimization**: Integrate functionality for resizing, compression, or format conversion to enhance performance and user experience.
// 2. **Secure Image Uploads**: Implement checks for image type and size, and consider scanning for malicious content to ensure uploads are safe.
// 3. **Storage & Delivery**: Utilize a CDN for efficient global delivery of images. Consider storing images in a service like AWS S3 for scalability.
// 4. **Accessibility Enhancements**: Utilize the `AltText` field to improve the accessibility of your application, providing descriptions for images that are informative for users who rely on screen readers.
// 5. **SEO Optimization**: Leverage the `AltText` and `Caption` fields to provide context-rich information to search engines, improving the discoverability of your images and listings.
// 6. **Image Versioning**: If images are updated frequently, implement a versioning system to track changes and allow for rollbacks to earlier versions if needed.
// 7. **Primary Image Functionality**: Use the `IsPrimary` flag to easily identify and display the main image for listings, enhancing the browsing experience.

// Implementing these features ensures your image model supports a scalable, accessible, and user-friendly application, with consideration for performance, security, and SEO.
