package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	BookingCreatedTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "booking_created_total",
			Help: "Total number of created bookings.",
		},
	)
	BookingConfirmedTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "booking_confirmed_total",
			Help: "Total number of confirmed bookings.",
		},
	)
	BookingExpiredTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "booking_expired_total",
			Help: "Total number of expired bookings.",
		},
	)
	BookingFailedTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "booking_failed_total",
			Help: "Total number of failed booking requests.",
		},
	)
)
