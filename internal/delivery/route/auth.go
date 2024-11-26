package route

import (
	authproto "banners_oto/gen/auth"
	sessionproto "banners_oto/gen/session"
	"banners_oto/gen/userproto"
	"banners_oto/internal/delivery/metrics"
	usauth "banners_oto/internal/usecase/auth"
	"banners_oto/internal/usecase/ucSession"
	ususer "banners_oto/internal/usecase/user"

	"banners_oto/microservices"
	"banners_oto/services"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func AddAuthRoute(mux *mux.Router, cluster *services.Cluster, clients *microservices.Clients, logger *zap.Logger, metrics *metrics.Metrics) {

	usManagerClient := userproto.NewUserManagerClient(clients.UserConn)
	usecaseUser := ususer.NewUsecaseLayer(usManagerClient, metrics)

	grpsSessionClient := sessionproto.NewSessionManagerClient(clients.SessionConn)
	usecaseSession := ucSession.NewUsecaseLayer(grpsSessionClient, metrics)

	grpcAutnClient := authproto.NewAuthManagerClient(clients.AuthConn)
	usecaseAuth := usauth.NewUsecaseLayer(grpcAutnClient, metrics)

}
