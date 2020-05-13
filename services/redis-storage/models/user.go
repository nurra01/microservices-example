package models

// RegisterUser defines type for user signup
type RegisterUser struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Verified  bool   `json:"verified"`
}
