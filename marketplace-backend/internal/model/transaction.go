package model

import (
	"gorm.io/gorm"
	"time"
)

// TransactionStatus defines possible states of a transaction
type TransactionStatus string

const (
	TransactionPending   TransactionStatus = "pending"
	TransactionCompleted TransactionStatus = "completed"
	TransactionCancelled TransactionStatus = "cancelled"
)

// Transaction represents a financial transaction between users for a listing.
type Transaction struct {
	gorm.Model
	BuyerID     uint              `gorm:"not null"` // ID of the buyer
	SellerID    uint              `gorm:"not null"` // ID of the seller
	ListingID   uint              `gorm:"not null"` // Associated listing
	Status      TransactionStatus `gorm:"type:varchar(20);not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
