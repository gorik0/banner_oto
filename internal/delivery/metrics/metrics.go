package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	TotalNumberOfRequests    prometheus.Counter
	RequestTime              *prometheus.HistogramVec
	NumberOfSpecificRequests *prometheus.CounterVec
	MicroserviceTimeout      *prometheus.HistogramVec
	MicroserviceErrors       *prometheus.CounterVec
	DatabaseDuration         *prometheus.HistogramVec
	PopularRestaurant        *prometheus.CounterVec
	PopularFood              *prometheus.CounterVec
	PopularCategory          *prometheus.CounterVec
}

func NewMetrics(namespace string, reg prometheus.Registerer) *Metrics {

}
