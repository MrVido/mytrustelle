package model

import (
	"gorm.io/gorm"
)

// ReportType defines categories for reports
type ReportType string

const (
	ReportTypeListing ReportType = "listing"
	ReportTypeUser    ReportType = "user"
	ReportTypeReview  ReportType = "review" // Added to allow reporting reviews
)

// ReportStatus defines the possible states of a report's resolution process
type ReportStatus string

const (
	ReportStatusPending   ReportStatus = "pending"
	ReportStatusReviewed  ReportStatus = "reviewed"
	ReportStatusResolved  ReportStatus = "resolved"
	ReportStatusDismissed ReportStatus = "dismissed"
)

// Report represents a report made by a user against a listing, another user, or a review.
type Report struct {
	gorm.Model
	ReporterID uint         `gorm:"index;not null"` // Indexed for quicker lookups by ReporterID
	ReportType ReportType   `gorm:"type:varchar(20);index"` // Indexed for filtering by report type
	TargetID   uint         `gorm:"index;not null"` // Indexed for quicker lookups by TargetID
	Reason     string       `gorm:"type:text;not null"` // Reason for the report
	Details    string       `gorm:"type:text"` // Additional details about the report
	Status     ReportStatus `gorm:"type:varchar(20);default:'pending';index"` // Current status of the report
	ResolvedAt *time.Time   // Time when the report was resolved
	ResolverID uint         // User ID of the admin who resolved the report (if applicable)
	ResolutionDetails string `gorm:"type:text"` // Explanation or details of how the report was resolved
}

// Key Enhancements:

// 1. **Report on Reviews**: Adding `ReportTypeReview` allows users to report problematic reviews, enhancing moderation capabilities.

// 2. **Report Status and Workflow**: The `Status` field with defined statuses (`pending`, `reviewed`, `resolved`, `dismissed`) facilitates a clear workflow for handling reports.

// 3. **Resolution Tracking**: `ResolvedAt`, `ResolverID`, and `ResolutionDetails` provide a detailed track of how and when a report was resolved, improving transparency and accountability.

// 4. **Indexed Fields**: Indexing `ReporterID`, `ReportType`, `TargetID`, and `Status` optimizes query performance, especially for administrative
