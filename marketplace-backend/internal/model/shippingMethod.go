package model

import (
	"gorm.io/gorm"
	"time"
)

// ShippingMethod represents various shipping options available for items listed on the marketplace.
type ShippingMethod struct {
	gorm.Model
	Name               string    `gorm:"not null"` // The name of the shipping method, e.g., "Standard", "Express".
	Description        string    `gorm:"type:text"` // A short description of the shipping method.
	Cost               float64   `gorm:"not null"` // The cost associated with this shipping method.
	EstimatedDelivery  string    `gorm:"not null"` // Estimated delivery time (e.g., "3-5 business days").
	IsActive           bool      `gorm:"default:true"` // Indicates if the shipping method is active and should be displayed as an option.
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

// Consider adding relationships to other models if necessary, for example, linking ShippingMethods to specific Listings or Orders.
// Key Points:

// Flexibility and Scalability: This model is designed to be flexible, allowing for easy addition or removal of shipping methods as the business scales or as logistics needs evolve.
// Active Status Management: The IsActive field allows for easy management of which shipping options are currently available, helping to manage seasonal shipping options or temporarily unavailable methods without removing them from the system.
// User and Seller Convenience: By providing clear names, costs, and estimated delivery times, this model aims to enhance the shopping experience for users and enable sellers to choose appropriate shipping methods for their listings.
// This model forms the basis for managing shipping options within your marketplace, ensuring that sellers can offer a variety of shipping choices to meet buyer preferences and logistical requirements. Next steps include integrating this model with listings and orders to apply shipping methods to transactions dynamically.