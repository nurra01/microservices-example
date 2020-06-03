package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

// pointer to the db connection to use in other functions
var dbConn *sqlx.DB

// Connect connects to the database and returns an instance
func Connect() error {
	// get all environment variables for DB connection
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	driver := os.Getenv("DB_DRIVER")

	// if missing env variables, use default
	if host == "" || port == "" || user == "" || pass == "" || dbName == "" || driver == "" {
		driver = "postgres"
		host = "database"
		port = "5432"
		dbName = "kafka-example"
		user = "admin"
		pass = "admin"
	}

	// DB URI "postgres://username:passw@host:5432/dbName?sslmode=disable"
	connURI := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		driver, user, pass, host, port, dbName)

	// connect to the DB
	conn, err := sqlx.Connect(driver, connURI)
	if err != nil {
		return fmt.Errorf("failed to connect DB. Error: %s", err.Error())
	}

	// ping DB to be sure
	err = conn.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping DB. Error: %s", err.Error())
	}

	dbConn = conn // set pointer to global dbConn
	return nil
}
