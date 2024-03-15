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
        Email    string `json:"email"` // Capture customer email for receipt.
        UserID   uint   `json:"userId"` // Example: User ID making the payment.
    }

    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment request"})
        return
    }

    // Enhanced validations
    if request.Amount <= 0 || request.Currency == "" || request.UserID == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request parameters"})
        return
    }

    // Optionally, validate user ID against the database to ensure it's a valid user
    if !util.ValidateUserID(request.UserID) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
        return
    }

    // Additional parameters like customer ID can be used if you're managing Stripe Customer objects.
    params := &stripe.PaymentIntentParams{
        Amount:   stripe.Int64(request.Amount),
        Currency: stripe.String(request.Currency),
        ReceiptEmail: stripe.String(request.Email),
        Metadata: map[string]string{
            "userId": strconv.Itoa(int(request.UserID)), // Keep track of which user made the payment
        },
        // Example: Using a Stripe Customer ID for better management of recurring customers.
        // Customer: stripe.String("cust_123"),
    }

    paymentIntent, err := util.CreatePaymentIntent(params) // Assume util.CreatePaymentIntent abstracts the Stripe call.
    if err != nil {
        handleStripeError(err, c) // A function to categorize Stripe errors and respond appropriately.
        return
    }

    c.JSON(http.StatusOK, gin.H{"clientSecret": paymentIntent.ClientSecret})
}

// handleStripeError categorizes Stripe errors and responds appropriately.
func handleStripeError(err error, c *gin.Context) {
    if stripeErr, ok := err.(*stripe.Error); ok {
        // Switch on `stripeErr.Code` for specific Stripe error codes you wish to handle.
        switch stripeErr.Code {
        case stripe.ErrorCodeCardDeclined:
            c.JSON(http.StatusPaymentRequired, gin.H{"error": "Card was declined"})
            return
        default:
            log.Printf("Stripe error occurred: %v", stripeErr)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred processing the payment"})
            return
        }
    } else {
        log.Printf("Failed to create payment intent: %v", err) // Log the error for internal tracking
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initiate payment"})
    }
}
