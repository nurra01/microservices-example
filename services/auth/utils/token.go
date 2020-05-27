package utils

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"services/auth/models"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateToken create a new access token
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

// ExtractAccessToken gets access token from request header
func ExtractAccessToken(req *http.Request) (string, error) {
	authHeader := req.Header.Get("Authorization")
	if authorizationHeader != "" {
		bearerToken := strings.Split(authorizationHeader, " ")
		if len(bearerToken) == 2 {
			return bearerToken[1], nil
		} else {
			return "", errors.New("Invalid authorization token")
		}
	} else {
		return "", errors.New("Missing 'Authorization' header")
	}
}

// TokenValid validates whether token valid
func TokenValid(bearerToken string) error {
	// parse bearer token
	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return err
	}

	// if token is invalid return message
	if !token.Valid {
		return errors.New("Invalid authorization token")
	}

	return nil
}
