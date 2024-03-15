package model

import (
	"gorm.io/gorm"
	"time"
)

// ExternalServiceLog logs external service interactions for debugging and auditing.
type ExternalServiceLog struct {
	gorm.Model
	ServiceName string    `gorm:"not null"` // Name of the external service (e.g., "Stripe", "SendGrid")
	Request     string    `gorm:"type:text"` // Details of the request made to the service
	Response    string    `gorm:"type:text"` // Service's response
	Status      string    `gorm:"not null"` // Status of the interaction (e.g., "success", "error")
	Timestamp   time.Time `gorm:"default:CURRENT_TIMESTAMP"` // When the interaction occurred
}

// Enhancements:
// - Troubleshooting support for integration issues.
// - Transparency in third-party service performance and reliability.
