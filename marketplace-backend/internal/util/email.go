package util

import (
    "fmt"
    "os"

    "github.com/sendgrid/sendgrid-go"
    "github.com/sendgrid/sendgrid-go/helpers/mail"
)

// getSendGridAPIKey safely retrieves the SendGrid API key from environment variables.
func getSendGridAPIKey() string {
    return os.Getenv("SENDGRID_API_KEY")
}

// SendEmail sends an email using SendGrid with enhanced error handling and logging.
func SendEmail(toEmail, subject, htmlContent string) error {
    from := mail.NewEmail("Your Platform Name", "noreply@yourplatform.com")
    to := mail.NewEmail("User", toEmail)
    message := mail.NewSingleEmail(from, subject, to, htmlContent, htmlContent)
    client := sendgrid.NewSendClient(getSendGridAPIKey())

    response, err := client.Send(message)
    if err != nil {
        fmt.Printf("Failed to send email: %v\n", err)
        return err
    }

    // Check the response status code for potential issues
    if response.StatusCode >= 400 {
        fmt.Printf("Error sending email, status code: %d, body: %s\n", response.StatusCode, response.Body)
        return fmt.Errorf("sendgrid returned non-200 status code: %d", response.StatusCode)
    }

    fmt.Printf("Email sent successfully, status code: %v\n", response.StatusCode)
    return nil
}
