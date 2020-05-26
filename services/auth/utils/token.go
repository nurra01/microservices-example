package utils

import (
	"fmt"
	"os"
	"services/auth/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateToken create a new token
func GenerateToken(user *models.User) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["firstName"] = user.FirstName
	claims["lastName"] = user.LastName
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix() // token expires in 30 minutes
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign a token. Error: %s", err.Error())
	}
	return signedToken, nil
}

// GenerateRefreshToken create a new refresh token with longer expiration time
func GenerateRefreshToken(user *models.User) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user"] = user
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix() // token expires in 7 days
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign a token. Error: %s", err.Error())
	}
	return signedToken, nil
}
