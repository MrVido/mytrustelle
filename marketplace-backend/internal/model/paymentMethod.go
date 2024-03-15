package model

import (
	"gorm.io/gorm"
	"time"
)

// PaymentMethodType defines possible types of payment methods
type PaymentMethodType string

const (
	CreditCard PaymentMethodType = "credit_card"
	PayPal     PaymentMethodType = "paypal"
	// Additional payment method types as needed
)

// PaymentMethod represents a user's payment method in the marketplace.
type PaymentMethod struct {
	gorm.Model
	UserID     uint                `gorm:"not null;index"` // Reference to the User model
	Type       PaymentMethodType   `gorm:"type:varchar(100);not null"` // Type of payment method
	Details    string              `gorm:"not null"` // Encrypted payment details
	IsDefault  bool                `gorm:"default:false"` // Indicates the default payment method
	IsActive   bool                `gorm:"default:true"` // Indicates if the payment method is active
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// Note: Store sensitive details like card numbers or PayPal accounts encrypted.
// Implement encryption/decryption logic in your service layer.

// Key Points:

// Type Safety and Clarity: By defining PaymentMethodType as a custom type, the model explicitly controls the types of payment methods supported, enhancing code readability and safety.
// Sensitive Data Handling: The Details field is intended to store encrypted data. Always ensure that sensitive information is encrypted before storage and decrypted only when necessary, following best security practices.
// User Experience: Fields like IsDefault and IsActive allow users to manage their payment methods more flexibly, setting a default for convenience and enabling/disabling methods as needed.
// This model sets the groundwork for handling transactions securely and efficiently, ensuring a seamless checkout process for users. Next, we would integrate this model with service logic to handle encryption, validation, and payment processing with external gateways.