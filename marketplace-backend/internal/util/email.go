package util

import (
    "context"
    "fmt"
    "os"

    "github.com/sendgrid/sendgrid-go"
    "github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendEmail sends an email using SendGrid.
func SendEmail(toEmail, subject, htmlContent string) error {
    from := mail.NewEmail("Your Platform Name", "noreply@yourplatform.com")
    to := mail.NewEmail("User", toEmail)
    message := mail.NewSingleEmail(from, subject, to, htmlContent, htmlContent)
    client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

    response, err := client.Send(message)
    if err != nil {
        return err
    } else {
        fmt.Printf("Email Sent: %v\n", response.StatusCode)
    }

    return nil
}
