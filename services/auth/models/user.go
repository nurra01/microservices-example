package models

import "golang.org/x/crypto/bcrypt"

// User defines registered and verified user
type User struct {
	ID        string `json:"-" db:"id"`
	FirstName string `json:"firstName" db:"first_name"`
	LastName  string `json:"lastName" db:"last_name"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"-" db:"password"`
	Verified  bool   `json:"-" db:"verified"`
}

// VerifyPassword compares user's hashed password with target password
func (u *User) VerifyPassword(targetPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(targetPassword))
}
