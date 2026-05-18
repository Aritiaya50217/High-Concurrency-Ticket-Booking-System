package kafka

import (
	"context"

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

func (c *Consumer) Consume(ctx context.Context, handler func([]byte) error) {
	for {
		message, err := c.reader.ReadMessage(ctx)
		if err != nil {
			continue
		}

		if err := handler(message.Value); err != nil {
			continue
		}

	}
}
