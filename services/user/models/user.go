package models

import "errors"

// RegisterUser defines type for user signup
type RegisterUser struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Verified  bool   `json:"verified"`
}

// Validate all fields to be right
func (u *RegisterUser) Validate() error {
	if u.FirstName == "" || u.LastName == "" || u.Email == "" {
		return errors.New("required body fields are missing")
	}
	return nil
}
