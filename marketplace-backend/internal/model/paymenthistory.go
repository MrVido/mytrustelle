package model

import (
    "gorm.io/gorm"
    "time"
)

// PaymentType categorizes the type of payment for better analysis and reporting
type PaymentType string

const (
    PaymentTypePurchase      PaymentType = "purchase"
    PaymentTypeSubscription  PaymentType = "subscription"
    // Add more as needed
)

// PaymentStatus represents the possible outcomes of a payment process
type PaymentStatus string

const (
    PaymentStatusCompleted  PaymentStatus = "completed"
    PaymentStatusPending    PaymentStatus = "pending"
    PaymentStatusFailed     PaymentStatus = "failed"
    // Add more as needed
)

// PaymentHistory records the details of financial transactions on the platform.
type PaymentHistory struct {
    gorm.Model
    UserID    uint            `gorm:"index;not null"` // User involved in the payment transaction
    Type      PaymentType     `gorm:"type:varchar(50); not null"` // Type of payment
    Amount    float64         `gorm:"not null"` // Amount transacted
    Currency  string          `gorm:"type:varchar(10); not null"` // Currency code (e.g., USD, EUR)
    Status    PaymentStatus   `gorm:"type:varchar(50); not null"` // Status of the payment
    Details   string          `gorm:"type:text"` // Additional details about the transaction
    Timestamp time.Time       `gorm:"index;not null"` // The time when the transaction occurred
}

// Key Enhancements:
// - **PaymentType and PaymentStatus Enums**: These provide clear categorization and status tracking for payments, which can be particularly useful for reports and analytics.
// - **Timestamp Indexing**: Enhances the performance of queries filtering payments by date, which is common in reporting and auditing.
// - **Details Field**: Offers flexibility to store additional information about the transaction, such as payment method used or reason for failure, which can be invaluable for customer support.

// This PaymentHistory model will serve as a comprehensive log for all transactions, aiding in financial management, user support, and strategic planning for your marketplace.
