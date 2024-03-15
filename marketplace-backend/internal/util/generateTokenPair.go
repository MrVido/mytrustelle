package util

import (
    "time"

    "github.com/dgrijalva/jwt-go"
)

// GenerateTokenPair generates both an access token and a refresh token for a user.
func GenerateTokenPair(userID uint) (accessToken string, refreshToken string, err error) {
    // Access Token generation (short-lived)
    accessToken, err = generateJWT(userID, 15*time.Minute) // Access token valid for 15 minutes
    if err != nil {
        return "", "", err
    }

    // Refresh Token generation (long-lived)
    refreshToken, err = generateJWT(userID, 24*time.Hour*7) // Refresh token valid for 7 days
    if err != nil {
        return "", "", err
    }

    return accessToken, refreshToken, nil
}

// generateJWT creates a signed JWT token with the given userID and validity duration.
func generateJWT(userID uint, duration time.Duration) (string, error) {
    expirationTime := time.Now().Add(duration)
    claims := &Claims{
        UserID: userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        return "", err
    }
    return tokenString, nil
}

// Claims struct to extend the standard jwt.Claims with the user ID
type Claims struct {
    UserID uint
    jwt.StandardClaims
}

var jwtSecret = []byte("your_secret_key") // Replace with your actual secret key
