package models

import (
	"errors"
	"regexp"
)

// RegisterUser defines type for user signup
type RegisterUser struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Verified  bool   `json:"verified"`
}

// Validate all fields to be right
func (u *RegisterUser) Validate() error {
	if u.FirstName == "" || u.LastName == "" || u.Email == "" || u.Password == "" {
		return errors.New("required body fields are missing")
	}

	// validate email to be valid
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !Re.MatchString(u.Email) {
		return errors.New("email is invalid")
	}

	// for simplicity just validate password to have at least 8 characters (don't do in production)
	if len(u.Password) < 8 {
		return errors.New("invalid password, password must be at least 8 characters long")
	}
	return nil
}
