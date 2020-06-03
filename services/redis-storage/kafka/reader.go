package kafka

import (
	"context"
	"fmt"
	"os"
	"services/redis-storage/redis"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

var reader *kafka.Reader // pointer to the reader, so Pull can use it

var brokerHost = os.Getenv("KAFKA_BROKER_HOST")
var brokerPort = os.Getenv("KAFKA_BROKER_PORT")
var regUserTopic = os.Getenv("KAFKA_REG_TOPIC") // topic for registered users
var verUserTopic = os.Getenv("KAFKA_VER_TOPIC") // topic for verified users
var clientID = os.Getenv("KAFKA_CLIENT_ID")

// ConfigureReader sets up a kafka reader
func ConfigureReader() (*kafka.Reader, error) {
	// if missing env variables, use default
	if brokerHost == "" || brokerPort == "" || regUserTopic == "" || clientID == "" {
		brokerHost = "kafka"
		brokerPort = "9092"
		regUserTopic = "register-user"
		clientID = "kafka-client-id"
	}

	brokerURL := fmt.Sprintf("%s:%s", brokerHost, brokerPort)
	brokers := []string{brokerURL}

	// config for reader
	config := kafka.ReaderConfig{
		Brokers:         brokers,
		GroupID:         clientID,
		Topic:           regUserTopic,
		MinBytes:        10e3,            // 10KB
		MaxBytes:        10e6,            // 10MB
		MaxWait:         1 * time.Second, // Maximum amount of time to wait for new data to come when fetching batches of messages from kafka.
		ReadLagInterval: -1,
	}
	// init a reader
	r := kafka.NewReader(config)
	reader = r
	return r, nil
}

// ReadMessages reads messages from kafka
func ReadMessages(log *logrus.Logger) {
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Error("error while receiving messages: %v", err.Error())
		}
		// when message received
		// save new user in redis with expiration time
		log.Info(string(msg.Value))

		verifyID := uuid.New().String()
		log.Info("Verify ID: ", verifyID)
		err = redis.SaveUser(verifyID, msg.Value)
		if err != nil {
			log.Error(err)
		}

		// push user to verify topic so he/she receives email with verification link
		err = Push(context.Background(), []byte(verifyID), msg.Value)
		if err != nil {
			log.Error(err)
		}
	}
}
