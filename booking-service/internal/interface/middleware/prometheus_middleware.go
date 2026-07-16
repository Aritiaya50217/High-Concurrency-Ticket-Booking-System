package middleware

import (
	"strconv"
	"time"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/infrastructure/metrics"
	"github.com/gin-gonic/gin"
)

func PrometheusMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		duration := time.Since(start).Seconds()
		status := strconv.Itoa(ctx.Writer.Status())

		metrics.HTTPRequestTotal.WithLabelValues(
			ctx.Request.Method, ctx.FullPath(), status,
		).Inc()

		metrics.HTTPRequestDuration.WithLabelValues(ctx.Request.Method, ctx.FullPath()).Observe(duration)
	}
}
