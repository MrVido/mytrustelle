package util

import (
	"fmt"
	"os"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

// Initialize Stripe with your secret key from an environment variable for better security.
func init() {
	stripeKey := os.Getenv("STRIPE_SECRET_KEY")
	if stripeKey == "" {
		panic("STRIPE_SECRET_KEY is not set.")
	}
	stripe.Key = stripeKey
}

// CreatePaymentIntent creates a new payment intent for a given amount with additional error handling and options.
func CreatePaymentIntent(amount int64, currency string) (*stripe.PaymentIntent, error) {
	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(amount),
		Currency:           stripe.String(currency),
		// Consider adding more parameters as needed, such as PaymentMethodTypes, Description, etc., to handle more complex payment scenarios.
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		return nil, fmt.Errorf("failed to create payment intent: %v", err)
	}

	return pi, nil
}
