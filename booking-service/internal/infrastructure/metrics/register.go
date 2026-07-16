package metrics

import "github.com/prometheus/client_golang/prometheus"

var Registery = prometheus.NewRegistry()

func Register() {
	Registery.MustRegister(BookingCreatedTotal, BookingConfirmedTotal, BookingExpiredTotal, BookingFailedTotal, HTTPRequestTotal, HTTPRequestDuration)
}
