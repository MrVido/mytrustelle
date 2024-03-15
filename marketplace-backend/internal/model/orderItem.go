package model

import (
	"gorm.io/gorm"
)

// OrderItem represents a specific item purchased within an order.
type OrderItem struct {
	gorm.Model
	OrderID       uint    `gorm:"not null;index"` // ID of the associated order
	ListingID     uint    `gorm:"not null"` // ID of the listing being purchased
	Quantity      int     `gorm:"not null"` // Quantity of the item
	PricePerItem  float64 `gorm:"not null"` // Price per item at the time of purchase
	Subtotal      float64 `gorm:"not null"` // Subtotal price for this item (Quantity * PricePerItem)
	Customization string  `gorm:"type:text"` // Customization details, if applicable
}

// Enhancements could include:
// - Adding a reference to a Product or InventoryItem model for platforms managing stock directly.
// - Customization details for products that offer personalization options.
// - Automated subtotal calculation before saving to the database using GORM model hooks.
// Enhancements and Considerations:

// Direct Listing Reference: Linking directly to the ListingID allows easy access to product details, essential for order processing and customer inquiries.

// Customization Details: The Customization field caters to personalized items, enabling sellers to capture specific buyer preferences at the time of order.

// Automated Subtotal Calculation: Implement GORM's BeforeSave hook to automatically calculate the Subtotal based on Quantity and PricePerItem, ensuring data consistency and reducing manual calculation errors.

// Inventory Management Integration: For platforms managing stock levels, linking OrderItem to an InventoryItem model could automate stock adjustments post-purchase, enhancing inventory accuracy.

// Scalability and Flexibility: This structured yet flexible model supports a wide range of e-commerce scenarios, from simple orders to complex, customized purchases, providing a solid foundation for order item management.

// This dedicated OrderItem model enhances the order management system's robustness, supporting detailed tracking and management of individual items within orders. It lays the groundwork for sophisticated order processing capabilities, including support for product customizations and integration with inventory management systems.