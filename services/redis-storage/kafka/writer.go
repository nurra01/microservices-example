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

// use writer pointer so we can push messages in Push
var writer *kafka.Writer

// ConfigureWriter sets up a kafka writer
func ConfigureWriter() (w *kafka.Writer, err error) {
	brokerURL := fmt.Sprintf("%s:%s", os.Getenv("KAFKA_BROKER_HOST"), os.Getenv("KAFKA_BROKER_PORT"))
	brokers := []string{brokerURL}
	topic := os.Getenv("KAFKA_VER_TOPIC")
	clientID := os.Getenv("KAFKA_CLIENT_ID")

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
	w = kafka.NewWriter(config)
	writer = w
	return w, nil
}

// Push writes a new message
func Push(parent context.Context, key, value []byte) error {
	// message to write to kafka
	msg := kafka.Message{
		Key:   key,
		Value: value,
		Time:  time.Now(),
	}
	return writer.WriteMessages(parent, msg)
}
