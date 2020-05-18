package models

import "errors"

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
	return nil
}
