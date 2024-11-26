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

	AddAuthRoute(mux, cluster, clients, logger, m)
	AddUserRouter(mux, cluster, clients, logger, m)
	AddRestRouter(mux, cluster, clients, logger, m)
	AddOrderRouter(mux, cluster, clients, logger, m)
	AddQuizRouter(mux, cluster, clients, logger, m)
	AddPaymentRouter(mux, cluster, clients, logger, m)

	handler := AddMiddleware(mux, cluster, clients, logger, m)
	logger.Info("end of handlers definition")
	return handler
}
