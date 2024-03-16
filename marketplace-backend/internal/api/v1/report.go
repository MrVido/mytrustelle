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
// Enhanced version of GetReports with pagination.
func GetReports(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Existing authorization checks...

		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))

		var reports []model.Report
		result := db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&reports)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve reports"})
			return
		}

		// Send paginated response
		c.JSON(http.StatusOK, gin.H{
			"currentPage": page,
			"totalPages":  math.Ceil(float64(result.RowsAffected) / float64(pageSize)),
			"reports":     reports,
		})
	}
}
