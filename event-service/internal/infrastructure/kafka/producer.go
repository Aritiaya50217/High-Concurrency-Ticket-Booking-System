package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
	kafkago "github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(brokers []string, topic string) *Producer {
	writer := &kafkago.Writer{
		Addr:     kafkago.TCP(brokers...),
		Topic:    topic,
		Balancer: &kafkago.LeastBytes{},
	}
	return &Producer{writer: writer}
}

func (p *Producer) Publish(ctx context.Context, topic string, data any) error {

	log.Println("Publish to kafka :", data)

	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return p.writer.WriteMessages(ctx, kafkago.Message{
		Topic: topic,
		Value: payload,
	})
}
