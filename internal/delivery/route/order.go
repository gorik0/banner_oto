package route

import (
	restProto "banners_oto/gen/rest"
	userProto "banners_oto/gen/user"
	"banners_oto/internal/delivery/metrics"
	orderdelivery "banners_oto/internal/delivery/order"
	"banners_oto/internal/repository/food"
	"banners_oto/internal/repository/order"
	orderUsec "banners_oto/internal/usecase/order"
	"banners_oto/microservices"
	"banners_oto/services"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func AddOrderRouter(mux *mux.Router, cluster *services.Cluster, clients *microservices.Clients, logger *zap.Logger, metrics *metrics.Metrics) {
	repoOrder := order.NewRepoLayer(cluster.PsqlClient, metrics)
	repoFood := food.NewRepo(cluster.PsqlClient, metrics)
	//init user grpc client
	grpcUserClient := userProto.NewUserManagerClient(clients.UserConn)
	//init rest grpc client
	grpcRestClient := restProto.NewRestWorkerClient(clients.RestConn)

	usecaseOrder := orderUsec.NewUsecaseLayer(repoOrder, repoFood, grpcUserClient, grpcRestClient, metrics)
	handler := orderdelivery.NewOrderHandler(usecaseOrder, logger)

	mux.HandleFunc("/api/v1/order", handler.GetBasket).Methods("GET")

	mux.HandleFunc("/api/v1/order/{id}", handler.GetOrderById).Methods("GET")

	mux.HandleFunc("/api/v1/promocode", handler.SetPromocode).Methods("POST")
	mux.HandleFunc("/api/v1/promocode", handler.GetAllPromocode).Methods("GET")

	mux.HandleFunc("/api/v1/orders/current", handler.GetCurrentOrders).Methods("GET")
	mux.HandleFunc("/api/v1/orders/archive", handler.GetArchiveOrders).Methods("GET")

	mux.HandleFunc("/api/v1/order/update_address", handler.UpdateAddress).Methods("PUT")
	mux.HandleFunc("/api/v1/order/pay", handler.Pay).Methods("PUT")
	mux.HandleFunc("/api/v1/order/clean", handler.Clean).Methods("DELETE")

	mux.HandleFunc("/api/v1/order/food/add", handler.AddFood).Methods("POST")
	mux.HandleFunc("/api/v1/order/food/update_count", handler.UpdateFoodCount).Methods("PUT")
	mux.HandleFunc("/api/v1/order/food/delete/{food_id}", handler.DeleteFoodFromOrder).Methods("DELETE")
}
