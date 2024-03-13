package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	// Consider importing an error reporting service if you use one, e.g., Sentry, Rollbar.
)

// Recovery is a middleware that recovers from any panics and writes a 500 if there was one.
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Capture the error stack trace for more detailed logging.
				stackTrace := debug.Stack()

				// Log the error and stack trace. Consider logging to a file or a logging system as well.
				logMessage := fmt.Sprintf("Panic recovered: %v\n%s", err, stackTrace)
				fmt.Println(logMessage) // Example logging, replace with log.Printf, log to file, or use a logging library.

				// Optionally, report the error and stack trace to an external error tracking system (e.g., Sentry, Rollbar).
				// reportError(err, stackTrace)

				// Respond with a generic error message to avoid exposing sensitive error information.
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "Internal Server Error - please contact support",
				})
			}
		}()

		c.Next() // process request
	}
}

// Example function for reporting errors to an external system. Implement according to your chosen tool.
func reportError(err interface{}, stackTrace []byte) {
	// Example: Sending error report to an external system.
	// This function should be implemented according to the documentation of the error reporting service you're using.
	fmt.Printf("Reported error: %v with stack trace: %s\n", err, stackTrace)
	// Uncomment and implement when using an error reporting tool:
	// Sentry.CaptureException(fmt.Errorf("%v", err))
	// Rollbar.Critical(fmt.Errorf("%v", err), stackTrace)
}
