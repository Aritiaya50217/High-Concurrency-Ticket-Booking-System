package metrics

import "github.com/prometheus/client_golang/prometheus"

var Registry = prometheus.NewRegistry()

func Register() {
	Registry.MustRegister(SeatReservedTotal,
		SeatReleasedTotal,
		SeatReservationFailedTotal,

		HTTPRequestTotal,
		HTTPRequestDuration,

		KafkaMessagesConsumedTotal,
		KafkaMessagesProcessedTotal,
		KafkaConsumerErrorsTotal,
	)
}
