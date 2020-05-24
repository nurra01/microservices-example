package models

// User defines registered and verified user
type User struct {
	ID        string `json:"_" db:"id"`
	FirstName string `json:"firstName" db:"first_name"`
	LastName  string `json:"lastName" db:"last_name"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"_" db:"password"`
	Verified  bool   `json:"_" db:"verified"`
}
