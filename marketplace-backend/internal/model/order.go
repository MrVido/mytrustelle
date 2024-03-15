package model

import (
	"gorm.io/gorm"
	"time"
)

// OrderStatus defines possible states of an order
type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusShipped    OrderStatus = "shipped"
	OrderStatusDelivered  OrderStatus = "delivered"
	OrderStatusCancelled  OrderStatus = "cancelled"
	OrderStatusReturned   OrderStatus = "returned"
)

// Order represents a purchase made by a user on the marketplace.
type Order struct {
	gorm.Model
	UserID         uint            `gorm:"not null;index"` // ID of the user who made the order
	PaymentMethodID uint           `gorm:"not null"` // Chosen payment method for this order
	ShippingMethodID uint          `gorm:"not null"` // Chosen shipping method for this order
	TotalAmount    float64         `gorm:"not null"` // Total cost of the order
	Status         OrderStatus     `gorm:"type:varchar(20);not null;index"` // Current status of the order
	OrderItems     []OrderItem     `gorm:"foreignKey:OrderID"` // Items included in the order
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// OrderItem details each item included in an order.
type OrderItem struct {
	gorm.Model
	OrderID      uint    `gorm:"not null;index"` // Reference to the Order
	ListingID    uint    `gorm:"not null"` // ID of the listed item
	Quantity     int     `gorm:"not null"` // Quantity of the item ordered
	PriceAtOrder float64 `gorm:"not null"` // Price of the item at the time of the order
}


// Enhancements:

// Detailed Order Status: Including various statuses from pending to returned provides clear tracking of each order's lifecycle, enhancing customer support and post-purchase services.

// OrderItem Model: A separate model for OrderItem ensures flexibility in handling orders with multiple items, each potentially having different quantities and prices.

// Indexing: Indexing UserID and Status in the Order model, as well as OrderID in the OrderItem model, optimizes query performance, especially for fetching orders by user or status.

// Relationship to Payment and Shipping Methods: By referencing PaymentMethodID and ShippingMethodID, the model integrates closely with your platform's logistics and payment processing, providing a seamless transaction flow.

// Historical Price Tracking: Storing PriceAtOrder in OrderItem captures the price at the time of purchase, which is crucial for order history, analytics, and in case of returns or disputes.

// This model structure lays a comprehensive foundation for managing orders on your marketplace, from the point of sale through to delivery and potential returns, ensuring a robust framework for transaction management.





