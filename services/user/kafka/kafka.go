package kafka

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
)

// writer for pushing just registered users
var regUserWriter *kafka.Writer

// writer for pushing verified users
var verUserWriter *kafka.Writer

// Configure sets up a kafka writers
func Configure() (*kafka.Writer, *kafka.Writer, error) {
	brokerHost := os.Getenv("KAFKA_BROKER_HOST")
	brokerPort := os.Getenv("KAFKA_BROKER_PORT")
	regUserTopic := os.Getenv("KAFKA_REG_TOPIC") // topic for registered users
	verUserTopic := os.Getenv("KAFKA_VER_TOPIC") // topic for verified users
	clientID := os.Getenv("KAFKA_CLIENT_ID")

	// if missing env variables, use default
	if brokerHost == "" || brokerPort == "" || regUserTopic == "" || verUserTopic == "" || clientID == "" {
		brokerHost = "kafka"
		brokerPort = "9092"
		regUserTopic = "register-user"
		verUserTopic = "verified-user"
		clientID = "kafka-client-id"
	}

	brokerURL := fmt.Sprintf("%s:%s", brokerHost, brokerPort)
	brokers := []string{brokerURL}

	// configure writer for pushing registered users
	w1, err := configureWriter(regUserTopic, clientID, brokers)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to configure writer for registered users. %s", err.Error())
	}
	regUserWriter = w1

	// configure writer for pushing verified users
	w2, err := configureWriter(verUserTopic, clientID, brokers)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to configure writer for verified users. %s", err.Error())
	}
	verUserWriter = w2
	return w1, w2, nil
}

// configureWriter is abstract configure function
func configureWriter(topic string, clientID string, brokers []string) (*kafka.Writer, error) {
	if topic == "" || clientID == "" || len(brokers) < 1 {
		return nil, errors.New("failed to load required kafka env variables")
	}

	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
		ClientID: clientID,
	}

	config := kafka.WriterConfig{
		Brokers:          brokers,
		Topic:            topic,
		Balancer:         &kafka.LeastBytes{},
		Dialer:           dialer,
		WriteTimeout:     10 * time.Second,
		ReadTimeout:      10 * time.Second,
		CompressionCodec: snappy.NewCompressionCodec(),
	}
	w := kafka.NewWriter(config)
	return w, nil
}

// Push writes a new message to the kafka
func push(parent context.Context, writer *kafka.Writer, key, value []byte) error {
	// message to write to kafka
	msg := kafka.Message{
		Key:   key,
		Value: value,
		Time:  time.Now(),
	}
	return writer.WriteMessages(parent, msg)
}

// PushVerUser writes verified users
func PushVerUser(parent context.Context, key, value []byte) error {
	return push(parent, verUserWriter, key, value)
}

// PushRegUser writes registered users
func PushRegUser(parent context.Context, key, value []byte) error {
	return push(parent, regUserWriter, key, value)
}
