package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"services/user-storage/db"
	"services/user-storage/models"
	"time"

	"github.com/segmentio/kafka-go"
	_ "github.com/segmentio/kafka-go/snappy" // required to decode compressed messages
	"github.com/sirupsen/logrus"
)

var reader *kafka.Reader // pointer to the reader, so Pull can use it

// Configure sets up a kafka reader
func Configure() (*kafka.Reader, error) {
	brokerURL := fmt.Sprintf("%s:%s", os.Getenv("KAFKA_BROKER_HOST"), os.Getenv("KAFKA_BROKER_PORT"))
	brokers := []string{brokerURL}
	topic := os.Getenv("KAFKA_VER_TOPIC") // topic for verified users
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
		// when message received unmarshal it to object
		usr := &models.RegisterUser{}
		err = json.Unmarshal(msg.Value, usr)
		if err != nil {
			log.Error(err)
		}

		// additional check that user is verified
		if usr.Verified {
			log.Info(usr)
			// save new user to the db
			err = db.SaveNewUser(usr)
			if err != nil {
				log.Error(err)
			}
		}
	}
}
