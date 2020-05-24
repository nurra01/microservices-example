package models

// RegisterUser defines type for user signup
type RegisterUser struct {
	ID        string `json:"id" db:"id"`
	FirstName string `json:"firstName" db:"first_name"`
	LastName  string `json:"lastName" db:"last_name"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	Verified  bool   `json:"verified" db:"verified"`
}
