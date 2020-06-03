package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
)

// use writer pointer so we can push messages in Push
var writer *kafka.Writer

// ConfigureWriter sets up a kafka writer to produce messages to 'verifiy user' topic
func ConfigureWriter() (w *kafka.Writer, err error) {
	// if missing env variables, use default
	if brokerHost == "" || brokerPort == "" || verUserTopic == "" || clientID == "" {
		brokerHost = "kafka"
		brokerPort = "9092"
		verUserTopic = "verify-user"
		clientID = "kafka-client-id"
	}

	brokerURL := fmt.Sprintf("%s:%s", brokerHost, brokerPort)
	brokers := []string{brokerURL}

	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
		ClientID: clientID,
	}

	config := kafka.WriterConfig{
		Brokers:          brokers,
		Topic:            verUserTopic,
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
