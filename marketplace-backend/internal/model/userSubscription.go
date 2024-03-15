package model

import (
	"gorm.io/gorm"
	"time"
)

// UserSubscription enhances the management of user subscriptions, adding next payment tracking.
type UserSubscription struct {
	gorm.Model
	UserID             uint      `gorm:"not null;index"` // Identifies the user
	SubscriptionPlanID uint      `gorm:"not null;index"` // Identifies the subscription plan
	StartDate          time.Time `gorm:"not null"` // When the subscription begins
	EndDate            time.Time `gorm:"not null"` // When the subscription is set to end
	IsActive           bool      `gorm:"default:true"` // Current active status of the subscription
	NextPaymentDate    time.Time `gorm:"not null"` // Date of the next payment
	AutoRenew          bool      `gorm:"default:false"` // Whether the subscription auto-renews
	CancellationDate   *time.Time // Date of cancellation, if applicable
}

// Including NextPaymentDate serves several purposes:
// - **Clarity for Users**: Provides users with precise information on when their next subscription payment will occur, enhancing transparency.
// - **Financial Planning**: Assists in forecasting upcoming payments and managing cash flow for both users and the platform.
// - **Subscription Lifecycle Management**: Facilitates automated processes for renewals, reminders, and possibly triggering early renewal incentives or grace periods for payment failures.

// This adjustment makes the subscription management system more comprehensive, aiding in better financial planning and user communication.
