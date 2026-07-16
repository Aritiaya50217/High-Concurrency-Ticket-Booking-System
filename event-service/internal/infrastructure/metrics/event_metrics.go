package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	SeatReservedTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "seat_reserved_total",
			Help: "Total number of successfully reserved seats.",
		},
	)

	SeatReleasedTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "seat_released_total",
			Help: "Total number of released seats. ",
		},
	)

	SeatReservationFailedTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "seat_reservation_failed_total",
			Help: "Total number of failed seat reservations.",
		},
	)
)
