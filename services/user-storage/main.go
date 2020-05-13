package main

import (
	"services/user-storage/db"
	"services/user-storage/kafka"

	"services/user-storage/logger"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// init a logger
	log := logger.NewLogger()

	// load .env file
	err := godotenv.Load("user-storage-service.env")
	if err != nil {
		log.Fatal("failed to load user-storage-service.env file")
	}

	// init a connection with DB
	_, err = db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	// setup required tables
	err = db.SetupTable()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("successfully setup required tables")

	// init a new kafka message reader
	reader, err := kafka.Configure()
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close() // close reader when exit from main

	// read messages
	kafka.ReadMessages(log)
}
