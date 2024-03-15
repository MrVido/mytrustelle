package model

import (
    "gorm.io/gorm"
    "time"
)

// AuditAction defines the type of action being audited
type AuditAction string

const (
    AuditActionCreate   AuditAction = "create"
    AuditActionUpdate   AuditAction = "update"
    AuditActionDelete   AuditAction = "delete"
    AuditActionLogin    AuditAction = "login"
    AuditActionLogout   AuditAction = "logout"
    // Extend with more actions as needed
)

// AuditLog captures user actions for auditing purposes.
type AuditLog struct {
    gorm.Model
    UserID    uint       `gorm:"index"` // Optional: not all actions may be user-initiated
    Action    AuditAction `gorm:"type:varchar(50); not null"`
    Detail    string     `gorm:"type:text"` // Detailed description of the action
    Timestamp time.Time  `gorm:"index;not null"` // Time when the action was taken
}

// Key Enhancements:
// - **Action Specificity**: The `AuditAction` type allows for categorizing logs more specifically, aiding in filtering and analysis.
// - **User Association**: Including a `UserID` helps in tracing actions back to specific users, crucial for security and accountability.
// - **Timestamp Indexing**: Facilitates efficient querying of logs by time, which is essential for reviewing recent activity or investigating incidents.
// - **Detail Richness**: The `Detail` field allows for storing comprehensive information about each action, enabling thorough audits and investigations.

// Implementing this `AuditLog` model provides a robust mechanism for monitoring and reviewing activities on your platform, which is vital for maintaining a secure and trustworthy environment.
