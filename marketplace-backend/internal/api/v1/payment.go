package v1

import (
	"marketplace-backend/internal/util"
	"net/http"
	"github.com/gin-gonic/gin"
)

// InitiatePayment handles the creation of a Stripe payment intent.
func InitiatePayment(c *gin.Context) {
	var request struct {
		Amount   int64  `json:"amount"`
		Currency string `json:"currency"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a payment intent through Stripe
	paymentIntent, err := util.CreatePaymentIntent(request.Amount, request.Currency)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment intent"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"clientSecret": paymentIntent.ClientSecret,
	})
}
