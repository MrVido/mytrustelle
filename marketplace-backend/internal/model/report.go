package model

import (
	"gorm.io/gorm"
	"time"
)

// ReportType defines categories for reports
type ReportType string

const (
	ReportTypeListing ReportType = "listing"
	ReportTypeUser    ReportType = "user"
)

// Report represents a report made by a user against a listing or another user.
type Report struct {
	gorm.Model
	ReporterID uint       `gorm:"not null"`           // ID of the user who made the report
	ReportType ReportType `gorm:"type:varchar(20);"`  // Type of report (listing or user)
	TargetID   uint       `gorm:"not null"`           // ID of the listing or user being reported
	Reason     string     `gorm:"type:text;not null"` // Reason for the report
	Details    string     `gorm:"type:text"`          // Additional details
	Resolved   bool       `gorm:"default:false"`      // Whether the report has been resolved
	CreatedAt  time.Time
}
