package db

import (
	"fmt"
	"services/user-storage/models"
)

// SaveNewUser saves a user in the DB
func SaveNewUser(user *models.RegisterUser) error {
	_, err := dbConn.NamedExec(`INSERT INTO users (id,firstName,lastName,email,verified) VALUES (:id,:firstName,:lastName,:email,:verified)`, user)
	if err != nil {
		return fmt.Errorf("failed to save a new user. Error: %v", err.Error())
	}
	return nil
}

// SetupTable creates users table
func SetupTable() error {
	_, err := dbConn.Exec(`CREATE TABLE IF NOT EXISTS users (
		id VARCHAR(100) NOT NULL PRIMARY KEY,
		firstName VARCHAR(100) NOT NULL,
		lastName VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL,
		verified BOOLEAN NOT NULL
	)`)
	if err != nil {
		return fmt.Errorf("failed to create a table. %s", err.Error())
	}
	return nil
}
