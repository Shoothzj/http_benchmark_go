package common

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	PromSummaryRequestLatency = promauto.NewSummary(prometheus.SummaryOpts{
		Name:       prometheus.BuildFQName("gin", "request", "latency_ns"),
		Help:       "gin_request_latency_ns",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	})
)
