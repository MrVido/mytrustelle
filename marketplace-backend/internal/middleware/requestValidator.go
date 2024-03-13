package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// requestValidator uses the validator package to validate the request body.
var validate = validator.New()

// RequestValidator is a middleware for validating requests against provided struct schema.
func RequestValidator(schema interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bind the incoming JSON to the schema struct and check for errors
		if err := c.ShouldBindJSON(&schema); err != nil {
			// If there is a validation error, return an error response
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Perform struct validation
		if validationErr := validate.Struct(schema); validationErr != nil {
			// If struct validation fails, return an error response
			c.JSON(http.StatusBadRequest, gin.H{"validation_error": validationErr.Error()})
			c.Abort()
			return
		}

		// Continue down the chain if validation passes
		c.Next()
	}
}
