package main

import (
	"services/redis-storage/kafka"
	"services/redis-storage/logger"
	"services/redis-storage/redis"

	"github.com/joho/godotenv"
)

func main() {
	// init a logger
	log := logger.NewLogger()

	// load .env file
	err := godotenv.Load("redis-storage-service.env")
	if err != nil {
		log.Fatal("failed to load redis-storage-service.env file")
	}

	// setup redis client
	err = redis.ConnectRedis()
	if err != nil {
		log.Fatal(err)
	}

	// configure kafka writer and reader
	_, err = kafka.ConfigureWriter()
	if err != nil {
		log.Fatal(err)
	}

	_, err = kafka.ConfigureReader()
	if err != nil {
		log.Fatal(err)
	}

	// read messages
	kafka.ReadMessages(log)
}
