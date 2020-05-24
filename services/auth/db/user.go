package db

import (
	"services/auth/models"
)

// GetUser fetches user from DB by email
func GetUser(email string) (*models.User, error) {
	u := &models.User{}

	// postgres query
	query := `
		SELECT * FROM users WHERE email=$1 LIMIT 1
	`
	// get user from db by email
	err := dbConn.Get(u, query, email)
	if err != nil {
		return nil, err
	}

	return u, nil
}
