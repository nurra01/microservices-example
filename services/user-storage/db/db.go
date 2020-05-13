package db

import (
	"fmt"
	"os"
	"services/user-storage/models"

	"github.com/jmoiron/sqlx"
)

// pointer to the db connection to use in other functions
var dbConn *sqlx.DB

// ConnectDB connects to the database and returns an instance
func ConnectDB() (*sqlx.DB, error) {
	// get DB connection URI and driver name
	dbDriver, dbURI := getDbURI()
	// connect to the DB
	conn, err := sqlx.Connect(dbDriver, dbURI)
	if err != nil {
		return nil, fmt.Errorf("failed to connect DB. Error: %s", err.Error())
	}
	// ping DB to be sure
	err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping DB. Error: %s", err.Error())
	}
	dbConn = conn // set pointer to global dbConn
	return conn, nil
}

// returns DB connection URI
func getDbURI() (string, string) {
	// get all environment variables for DB connection
	dbConfig := &models.DBConfig{
		Host:   os.Getenv("DB_HOST"),
		Port:   os.Getenv("DB_PORT"),
		User:   os.Getenv("DB_USER"),
		Pass:   os.Getenv("DB_PASS"),
		Name:   os.Getenv("DB_NAME"),
		Driver: os.Getenv("DB_DRIVER"),
	}

	// DB URI "postgres://username:passw@host:5432/dbName?sslmode=disable"
	uri := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		dbConfig.Driver, dbConfig.User, dbConfig.Pass, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	return dbConfig.Driver, uri
}
