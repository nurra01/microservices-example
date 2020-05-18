package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes and salts a password
func HashPassword(pass string) (string, error) {
	// hash and salt password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash and salt password. %s", err.Error())
	}
	return string(hashedPass), nil
}
