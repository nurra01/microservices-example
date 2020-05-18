package kafka

import (
	"context"
	"errors"
	"fmt"
	"os"
	"services/redis-storage/redis"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

var reader *kafka.Reader // pointer to the reader, so Pull can use it

// ConfigureReader sets up a kafka reader
func ConfigureReader() (*kafka.Reader, error) {
	brokerURL := fmt.Sprintf("%s:%s", os.Getenv("KAFKA_BROKER_HOST"), os.Getenv("KAFKA_BROKER_PORT"))
	brokers := []string{brokerURL}
	topic := os.Getenv("KAFKA_REG_TOPIC")
	clientID := os.Getenv("KAFKA_CLIENT_ID")

	if topic == "" || clientID == "" || len(brokers) < 1 {
		return nil, errors.New("failed to load required kafka env variables")
	}

	// config for reader
	config := kafka.ReaderConfig{
		Brokers:         brokers,
		GroupID:         clientID,
		Topic:           topic,
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

		// save token in redis with expiration time
		err = redis.SaveToken(msg.Value)
		if err != nil {
			log.Error(err)
		}
		log.Info("saved token in redis")

		// push user to verify topic so he/she receives email with verification link
		err = Push(context.Background(), []byte(verifyID), msg.Value)
		if err != nil {
			log.Error(err)
		}
	}
}
