package monitor

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

// Monitor structure holds all necessary data for monitoring
type Monitor struct {
	// Define necessary fields
	httpRequestsTotal *prometheus.CounterVec
}

// NewMonitor creates a new monitor
func NewMonitor() *Monitor {
	m := &Monitor{
		httpRequestsTotal: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_requests_total",
				Help: "Number of HTTP requests",
			},
			[]string{"path"},
		),
	}

	return m
}

// HandleMetrics exposes the Prometheus metrics endpoint
func (m *Monitor) HandleMetrics() http.Handler {
	return promhttp.Handler()
}

// TrackRequest tracks an HTTP request
func (m *Monitor) TrackRequest(path string) {
	m.httpRequestsTotal.With(prometheus.Labels{"path": path}).Inc()
}
