package middleware

import (
	"strconv"
	"time"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/event-service/internal/infrastructure/metrics"
	"github.com/gin-gonic/gin"
)

func PrometheusMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		path := ctx.FullPath()
		if path == "" {
			path = ctx.Request.URL.Path
		}

		metrics.HTTPRequestTotal.WithLabelValues(ctx.Request.Method, path, strconv.Itoa(ctx.Writer.Status())).Inc()

		metrics.HTTPRequestDuration.WithLabelValues(ctx.Request.Method, path).Observe(float64(time.Since(start).Seconds()))
	}
}
