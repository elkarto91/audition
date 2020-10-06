package models

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"time"
)

func RecordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(10 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "audition_processed_ops_total",
		Help: "The total number of processed events",
	})
)
