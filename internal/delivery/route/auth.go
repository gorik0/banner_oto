package route

import (
	authproto "banners_oto/gen/auth"
	sessionproto "banners_oto/gen/session"
	"banners_oto/gen/user"
	dAuth "banners_oto/internal/delivery/auth"
	"banners_oto/internal/delivery/metrics"
	usauth "banners_oto/internal/usecase/auth"
	"banners_oto/internal/usecase/session"
	ususer "banners_oto/internal/usecase/user"

	"banners_oto/microservices"
	"banners_oto/services"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func AddAuthRoute(mux *mux.Router, cluster *services.Cluster, clients *microservices.Clients, logger *zap.Logger, metrics *metrics.Metrics) {

	usManagerClient := user.NewUserManagerClient(clients.UserConn)
	usecaseUser := ususer.NewUsecaseLayer(usManagerClient, metrics)

	grpsSessionClient := sessionproto.NewSessionManagerClient(clients.SessionConn)
	usecaseSession := session.NewUsecaseLayer(grpsSessionClient, metrics)

	grpcAutnClient := authproto.NewAuthManagerClient(clients.AuthConn)
	usecaseAuth := usauth.NewUsecaseLayer(grpcAutnClient, metrics)

	deliveryAuth := dAuth.NewDeliveryLayer(usecaseSession, usecaseAuth, usecaseUser, logger, metrics)

	mux.HandleFunc("/api/v1/vk", deliveryAuth.AuthVk).Methods("POST").Name("vk-auth")
	mux.HandleFunc("/api/v1/signin", deliveryAuth.SignIn).Methods("POST").Name("signin")
	mux.HandleFunc("/api/v1/signup", deliveryAuth.SignUp).Methods("POST").Name("signup")
	mux.HandleFunc("/api/v1/signout", deliveryAuth.SignOut).Methods("POST").Name("signout")
}
