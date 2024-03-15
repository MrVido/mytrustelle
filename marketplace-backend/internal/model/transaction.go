package model

import (
	"gorm.io/gorm"
)

// TransactionStatus defines possible states of a transaction.
type TransactionStatus string

const (
	TransactionPending   TransactionStatus = "pending"
	TransactionCompleted TransactionStatus = "completed"
	TransactionFailed    TransactionStatus = "failed"
	TransactionCancelled TransactionStatus = "cancelled"
)

// Transaction represents a financial transaction between users for a listing.
type Transaction struct {
	gorm.Model
	BuyerID     uint              `gorm:"index;not null"` // Index for faster queries by BuyerID
	SellerID    uint              `gorm:"index;not null"` // Index for faster queries by SellerID
	ListingID   uint              `gorm:"index;not null"` // Index for association with a specific listing
	Status      TransactionStatus `gorm:"type:varchar(20);not null;index"` // Indexed for querying by status
	Amount      float64           `gorm:"not null"`       // Transaction amount
	Currency    string            `gorm:"type:varchar(3);not null"` // ISO currency code, e.g., USD, EUR
	PaymentMethod string          `gorm:"type:varchar(50);not null"` // Payment method used (e.g., PayPal, Credit Card, Stripe)
	PaymentGatewayTransactionID string `gorm:"uniqueIndex"` // Unique identifier for the transaction from the payment gateway
	SecurityCode string          `gorm:"-"` // Security code for transaction verification, not stored in DB
	Notes       string           `gorm:"type:text"`      // Optional notes about the transaction
	CompletedAt *time.Time       `gorm:"index"`          // Time when the transaction was completed, nil if pending or cancelled
}

// Key Enhancements:

// 1. **Enhanced Status Management**: Including a `TransactionFailed` status to accurately represent transactions that did not go through successfully.

// 2. **Transaction Amount and Currency**: Capturing both the amount and the currency used for the transaction to facilitate international transactions.

// 3. **Payment Method and Gateway ID**: Storing details about the payment method and a unique identifier from the payment gateway for reference, reconciliation, and potential refunds.

// 4. **Security and Verification**: The `SecurityCode` field (excluded from database storage with `gorm:"-"`) could be used transiently to verify transactions without storing sensitive information.

// 5. **Notes and Metadata**: Providing a `Notes` field for any additional information related to the transaction that might be useful for record-keeping.

// 6. **Completion Timestamp**: Tracking when a transaction was completed with `CompletedAt` allows for better management of transaction states and historical data analysis.

// 7. **Database Indexes**: Adding indexes on foreign keys and status fields improves performance for query operations, which is crucial for systems with a large number of transactions.

// These enhancements aim to create a comprehensive transaction model that supports detailed tracking, reporting, and analysis of financial transactions within your marketplace platform, accommodating a wide range of payment methods and facilitating international commerce.
