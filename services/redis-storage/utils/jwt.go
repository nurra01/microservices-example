package utils

import (
	"errors"
	"fmt"
	"os"
	"services/redis-storage/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CreateToken generates a new JWT token
func CreateToken(usr *models.RegisterUser) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return "", errors.New("missing 'SECRET_KEY' in environment variables")
	}
	// create JWT claims
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = usr.ID
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix() // expire time is 15 minutes

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign token with secret key
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign JWT token. %s", err.Error())
	}

	return signedToken, nil
}
