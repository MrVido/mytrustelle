package util

import (
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

// Initialize Stripe with your secret key.
func init() {
	stripe.Key = "your_stripe_secret_key"
}

// CreatePaymentIntent creates a new payment intent for a given amount.
func CreatePaymentIntent(amount int64, currency string) (*stripe.PaymentIntent, error) {
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String(currency),
	}
	return paymentintent.New(params)
}
