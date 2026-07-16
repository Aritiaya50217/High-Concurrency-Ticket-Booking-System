package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	UserRegisteredTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "user_registered_total",
			Help: "Total number of registered users.",
		},
	)

	UserLoginSuccessTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "user_login_success_total",
			Help: "Total number of successful logins.",
		},
		[]string{"result"},
	)

	UserLoginFailedTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "user_login_failed_total",
			Help: "Total number of failed logins.",
		},
		[]string{"result"},
	)

	UserGetProfileSuccessTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "user_get_profile_success_total",
			Help: "Total number of successful get profiles.",
		},
	)

	UserGetProfileFailedTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "user_get_profile_failed_total",
			Help: "Total number of failed get profiles.",
		},
	)
)
