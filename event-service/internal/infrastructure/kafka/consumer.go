package kafka

import (
	"context"
	"log"

	kafkago "github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafkago.Reader
}

func NewConsumer(brokers []string, topic string, groupID string) *Consumer {
	reader := kafkago.NewReader(kafkago.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: groupID,
	})
	return &Consumer{reader: reader}
}

func (c *Consumer) Start(ctx context.Context, handler func([]byte) error) {
	log.Println("event consumer started, waiting for messages...")

	for {
		message, err := c.reader.ReadMessage(ctx)
		if err != nil {
			log.Println("read kafka error:", err)
			continue
		}

		if err := handler(message.Value); err != nil {
			log.Println("handler error:", err)
			continue
		}

	}
}
