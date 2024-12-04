package route

import (
	protosession "banners_oto/gen/session"
	protouser "banners_oto/gen/user"
	"banners_oto/internal/delivery/metrics"
	"banners_oto/internal/delivery/user"
	ussession "banners_oto/internal/usecase/session"
	ususer "banners_oto/internal/usecase/user"
	"banners_oto/microservices"
	"banners_oto/services"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func AddUserRouter(mux *mux.Router, cluster *services.Cluster, clients *microservices.Clients, logger *zap.Logger, metrics *metrics.Metrics) {
	// init user grpc client
	grpcUserClient := protouser.NewUserManagerClient(clients.UserConn)
	usecaseUser := ususer.NewUsecaseLayer(grpcUserClient, metrics)
	// init session grpc client
	grpcSessionClient := protosession.NewSessionManagerClient(clients.SessionConn)
	usecaseSession := ussession.NewUsecaseLayer(grpcSessionClient, metrics)

	deliveryUser := user.NewDeliveryLayer(usecaseSession, usecaseUser, logger, metrics)

	mux.HandleFunc("/api/v1/user", deliveryUser.UserData).Methods("GET").Name("user_data")
	mux.HandleFunc("/api/v1/user", deliveryUser.UpdateInfo).Methods("PUT").Name("update_user")
	mux.HandleFunc("/api/v1/user/address", deliveryUser.UpdateAddress).Methods("PUT").Name("update_user_address")
	mux.HandleFunc("/api/v1/user/unauth_address", deliveryUser.UpdateUnauthAddress).Methods("PUT").Name("update_user_unauth_address")
	mux.HandleFunc("/api/v1/user/address", deliveryUser.UserAddress).Methods("GET").Name("user_address")
	mux.HandleFunc("/api/v1/user/new_password", deliveryUser.UpdatePassword).Methods("PUT").Name("update_user_password")
}
