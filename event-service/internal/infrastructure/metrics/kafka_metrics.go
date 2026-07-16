package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	KafkaMessagesConsumedTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "kafka_messages_consumed_total",
			Help: "Total number of Kafka messages consumed.",
		},
	)

	KafkaMessagesProcessedTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "kafka_messages_processed_total",
			Help: "Total number of Kafka messages processed successfully.",
		},
	)

	KafkaConsumerErrorsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "kafka_consumer_errors_total",
			Help: "Total number of Kafka consumer errors.",
		},
		[]string{"type"},
	)
)
