package util

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Assuming you've moved jwtSecret to use an environment variable for better security practices.
var jwtSecret = []byte(os.Getenv("JWT_SECRET")) // Ensure you set this environment variable.

// Claims struct will add custom claims which extend the standard jwt claims
type Claims struct {
	UserID uint
	jwt.StandardClaims
}

// GenerateToken generates a jwt token for a given user ID.
func GenerateToken(userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token expires after 24 hours
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)

	return tokenString, err
}

// ValidateToken validates the given token and returns the user ID.
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Handle the token expiration case
				return nil, fmt.Errorf("token has expired")
			} else {
				// Handle other validation errors (e.g., signature invalid)
				return nil, fmt.Errorf("invalid token: %v", err)
			}
		}
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
