package v1

import (
	"marketplace-backend/internal/model"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SubmitReport allows users to report listings or other users.
func SubmitReport(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var report model.Report
		if err := c.ShouldBindJSON(&report); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if result := db.Create(&report); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Report submitted successfully"})
	}
}

// GetReports retrieves all reports for administrators.
func GetReports(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reports []model.Report
		if result := db.Find(&reports); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, reports)
	}
}
