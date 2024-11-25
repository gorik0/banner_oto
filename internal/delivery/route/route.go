package route

import (
	"banners_oto/internal/delivery/metrics"
	"banners_oto/microservices"
	"banners_oto/services"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func Setup(mux *mux.Router, cluster *services.Cluster, clients *microservices.Clients, m *metrics.Metrics, reg *prometheus.Registry, logger *zap.Logger) http.Handler {
	logger.Info("Starting routing definition")

	// prometheus route
	promHttp := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	mux.Handle("/metrics", promHttp)

	// services route

}
