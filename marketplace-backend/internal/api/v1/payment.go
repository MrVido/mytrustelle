package v1

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/stripe/stripe-go/v72"
    "github.com/stripe/stripe-go/v72/paymentintent"
    "marketplace-backend/internal/util"
)

// Ensure Stripe's API key is set in your application's initialization phase.
// stripe.Key = "sk_test_yourSecretKey"

// InitiatePayment handles the creation of a Stripe payment intent with enhanced validations and parameters.
func InitiatePayment(c *gin.Context) {
    var request struct {
        Amount   int64  `json:"amount"`
        Currency string `json:"currency"`
        Email    string `json:"email"` // Optional: Capture customer email for receipt.
    }

    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment request"})
        return
    }

    // Validate request fields
    if request.Amount <= 0 || request.Currency == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount or currency"})
        return
    }

    // Create a payment intent through Stripe with additional parameters
    params := &stripe.PaymentIntentParams{
        Amount:   stripe.Int64(request.Amount),
        Currency: stripe.String(request.Currency),
        ReceiptEmail: stripe.String(request.Email), // Optional
        Metadata: map[string]string{
            "integration_check": "accept_a_payment",
        },
    }

    // Utilizing the utility package if it abstracts Stripe operations
    // Otherwise, directly call Stripe's API for creating a payment intent
    paymentIntent, err := paymentintent.New(params)
    if err != nil {
        log.Printf("Failed to create payment intent: %v", err) // Log the error for internal tracking
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initiate payment"})
        return
    }

    // Send the client secret to the client to complete the payment process
    c.JSON(http.StatusOK, gin.H{
        "clientSecret": paymentIntent.ClientSecret,
    })
}
