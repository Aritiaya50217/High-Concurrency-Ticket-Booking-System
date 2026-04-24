package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(brokers []string, topic, groupID string) *Consumer {
	return &Consumer{reader: kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: groupID,
	})}
}

func (c *Consumer) Start(ctx context.Context, handler func([]byte)) {
	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			log.Println("kafka error : ", err)
			continue
		}
		handler(msg.Value)
	}
}
