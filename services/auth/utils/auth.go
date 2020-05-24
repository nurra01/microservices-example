package utils

import "golang.org/x/crypto/bcrypt"

// VerifyPassword compares hashed password with target password
func VerifyPassword(hashedPass, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))
}
