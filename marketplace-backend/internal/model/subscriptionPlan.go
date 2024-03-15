package model

import (
	"gorm.io/gorm"
	"time"
)

// SubscriptionPlan defines the structure for subscription plans available on the platform.
type SubscriptionPlan struct {
	gorm.Model
	Name         string    `gorm:"unique;not null"` // Unique name of the subscription plan
	Description  string    `gorm:"type:text;not null"` // Description of what the plan offers
	Price        float64   `gorm:"not null"` // Monthly price of the subscription plan
	Duration     int       `gorm:"not null"` // Duration of the subscription in days
	Features     string    `gorm:"type:text;not null"` // Detailed list of features provided by the plan, could be JSON encoded for flexibility
	IsActive     bool      `gorm:"default:true"` // Whether the plan is currently offered
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// UserSubscription links users with their subscription plans and tracks the subscription status.
type UserSubscription struct {
	gorm.Model
	UserID              uint      `gorm:"not null;index"` // Reference to the User model
	SubscriptionPlanID  uint      `gorm:"not null;index"` // Reference to the SubscriptionPlan model
	StartDate           time.Time `gorm:"not null"` // When the subscription starts
	EndDate             time.Time `gorm:"not null"` // When the subscription ends
	IsActive            bool      `gorm:"default:true"` // Tracks if the subscription is active
	RenewalReminderSent bool      `gorm:"default:false"` // Indicates if a renewal reminder has been sent
}

// Key Enhancements:
// - **Features as JSON**: Storing the `Features` field as a JSON-encoded string (or leveraging a JSONB type in databases that support it) allows for storing a structured list of features that can easily be extended or modified without changing the database schema.
// - **Active Status Management**: `IsActive` in both models allows you to control the availability of plans and manage the active status of user subscriptions.
// - **Renewal Management**: `RenewalReminderSent` in `UserSubscription` helps manage communication with the user about upcoming renewals, improving the user experience and retention.

// These models provide a solid foundation for implementing a subscription feature in your marketplace, allowing for flexible management of subscription plans and user subscriptions.
