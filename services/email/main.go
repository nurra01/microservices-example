package main

import (
	"services/email/kafka"
	"services/email/logger"

	"github.com/joho/godotenv"
)

func main() {
	// init a logger
	log := logger.NewLogger()

	// load .env file
	err := godotenv.Load("email-service.env")
	if err != nil {
		log.Fatal("failed to load email-service.env file")
	}

	// init a new kafka message reader
	reader, err := kafka.Configure()
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	kafka.ReadMessages(log)
}
