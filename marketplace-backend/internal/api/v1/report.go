package v1

import (
	"marketplace-backend/internal/model"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SubmitReport allows users to report listings or other users, with added validation and authorization.
func SubmitReport(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var report model.Report
		if err := c.ShouldBindJSON(&report); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid report format"})
			return
		}

		// Add additional validation here if necessary, e.g., check if the reported entity exists.

		// Authorization check to ensure the user is logged in
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User must be logged in to submit a report"})
			return
		}

		// Set the reporter ID to the current user's ID
		report.ReporterID = userID.(uint)

		if result := db.Create(&report); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit report"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Report submitted successfully", "reportID": report.ID})
	}
}

// GetReports retrieves all reports for administrators, with authorization checks.
func GetReports(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authorization check to ensure the user is an admin
		userRole, exists := c.Get("userRole")
		if !exists || userRole != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access restricted to administrators"})
			return
		}

		var reports []model.Report
		if result := db.Find(&reports); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve reports"})
			return
		}

		c.JSON(http.StatusOK, reports)
	}
}
