package route

import (
	protorest "banners_oto/gen/rest"
	sessionproto "banners_oto/gen/session"
	protouser "banners_oto/gen/user"
	"banners_oto/internal/delivery/metrics"
	"banners_oto/internal/delivery/payment"
	"banners_oto/internal/repository/food"
	repoorder "banners_oto/internal/repository/order"
	usorder "banners_oto/internal/usecase/order"
	ussession "banners_oto/internal/usecase/session"
	"banners_oto/microservices"
	"banners_oto/services"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func AddPaymentRouter(mux *mux.Router, cluster *services.Cluster, clients *microservices.Clients, logger *zap.Logger, metrics *metrics.Metrics) {
	repoFood := food.NewRepo(cluster.PsqlClient, metrics)
	repoOrder := repoorder.NewRepoLayer(cluster.PsqlClient, metrics)

	// init session grpc client
	grpcSessionClient := sessionproto.NewSessionManagerClient(clients.SessionConn)
	usecaseSession := ussession.NewUsecaseLayer(grpcSessionClient, metrics)
	// init user grpc client
	grpcUserClient := protouser.NewUserManagerClient(clients.UserConn)
	// init rest grpc client
	grpcRestClient := protorest.NewRestWorkerClient(clients.RestConn)

	usecaseOrder := usorder.NewUsecaseLayer(repoOrder, repoFood, grpcUserClient, grpcRestClient, metrics)
	deliveryPayment := payment.NewPaymentDelivery(usecaseOrder, usecaseSession, logger, metrics)

	mux.HandleFunc("/api/v1/order/pay/url", deliveryPayment.OrderGetPayUrl).Methods("GET").Name("get-pay-url")
}
