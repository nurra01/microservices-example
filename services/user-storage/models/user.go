package models

// RegisterUser defines type for user signup
type RegisterUser struct {
	ID        string `json:"id" db:"id"`
	FirstName string `json:"firstName" db:"firstName"`
	LastName  string `json:"lastName" db:"lastName"`
	Email     string `json:"email" db:"email"`
	Verified  bool   `json:"verified" db:"verified"`
}
