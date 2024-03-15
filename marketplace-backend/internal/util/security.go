package util

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a bcrypt hash of the password using a specified cost.
// The cost determines how much time is needed to calculate a single bcrypt hash.
// Increasing the cost increases the security, but also the time to hash and verify passwords.
const passwordHashCost = 12 // This can be adjusted based on your security requirements and server capabilities.

// HashPassword generates a bcrypt hash of the password.
// It uses a custom cost defined above for a balance between security and performance.
func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), passwordHashCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedBytes), nil
}

// CheckPasswordHash compares a bcrypt hashed password with its possible plaintext equivalent.
// It returns true if the password and the hash match, false otherwise.
func CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			// The password does not match the hash
			return false, nil
		}
		// An unexpected error occurred
		return false, fmt.Errorf("unexpected error during hash comparison: %w", err)
	}
	return true, nil
}
