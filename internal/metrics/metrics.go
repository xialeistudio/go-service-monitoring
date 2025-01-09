package metrics

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	httpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request duration in seconds",
			Buckets: []float64{0.1, 0.3, 0.5, 0.7, 1, 3, 5, 7, 10},
		},
		[]string{"method", "path"},
	)
)

// Init initializes the metrics collector
func Init() {
}

// PrometheusMiddleware Gin middleware for collecting HTTP metrics
func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		// Skip /metrics path to avoid self-monitoring interference
		if c.Request.URL.Path != "/metrics" {
			duration := time.Since(start).Seconds()
			status := c.Writer.Status()

			httpRequestsTotal.WithLabelValues(
				c.Request.Method,
				c.Request.URL.Path,
				string(rune(status)),
			).Inc()

			httpRequestDuration.WithLabelValues(
				c.Request.Method,
				c.Request.URL.Path,
			).Observe(duration)
		}
	}
}
