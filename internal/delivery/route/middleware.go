package route

import (
	"banners_oto/config"
	protosession "banners_oto/gen/session"
	protouser "banners_oto/gen/user"
	"banners_oto/internal/delivery/metrics"
	http_middleware "banners_oto/internal/middleware/http"
	ussession "banners_oto/internal/usecase/session"
	ususer "banners_oto/internal/usecase/user"
	"banners_oto/microservices"
	"banners_oto/services"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func AddMiddleware(mux *mux.Router, cluster *services.Cluster, clients *microservices.Clients, logger *zap.Logger, m *metrics.Metrics) http.Handler {
	// init user microservice client
	grpcUserClient := protouser.NewUserManagerClient(clients.UserConn)
	usecaseUser := ususer.NewUsecaseLayer(grpcUserClient, m)
	// init session microservice client
	grpcSessionClient := protosession.NewSessionManagerClient(clients.SessionConn)
	usecaseSession := ussession.NewUsecaseLayer(grpcSessionClient, m)

	handler := http_middleware.SessionAuthentication(mux, usecaseUser, usecaseSession, &config.Config.Redis, logger)
	handler = http_middleware.Csrf(handler, usecaseSession, &config.Config, logger, m)
	handler = http_middleware.Cors(handler, logger)
	handler = http_middleware.Access(handler, logger)
	handler = http_middleware.Metrics(handler, m)

	return handler
}
