package config

import (
	"time"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

// ConfigKafka kafka
func ConfigKafka(kafkaBrokerUrls []string, clientID string, topic string) (w *kafka.Writer, err error) {
	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
		ClientID: clientID,
	}

	config := kafka.WriterConfig{
		Brokers:      kafkaBrokerUrls,
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		Dialer:       dialer,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	w = kafka.NewWriter(config)
	writer = w
	return w, nil
}
