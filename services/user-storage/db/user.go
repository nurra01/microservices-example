package db

import (
	"fmt"
	"services/user-storage/models"
)

// SaveNewUser saves a user in the DB
func SaveNewUser(user *models.RegisterUser) error {
	_, err := dbConn.NamedExec(`INSERT INTO users (id,first_name,last_name,email,password,verified) VALUES (:id,:first_name,:last_name,:email,:password,:verified)`, user)
	if err != nil {
		return fmt.Errorf("failed to save a new user. Error: %v", err.Error())
	}
	return nil
}

// SetupTable creates users table
func SetupTable() error {
	_, err := dbConn.Exec(`CREATE TABLE IF NOT EXISTS users (
		id VARCHAR(100) NOT NULL PRIMARY KEY,
		first_name VARCHAR(100) NOT NULL,
		last_name VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL,
		password VARCHAR(255) NOT NULL,
		verified BOOLEAN NOT NULL
	)`)
	if err != nil {
		return fmt.Errorf("failed to create a table. %s", err.Error())
	}
	return nil
}
