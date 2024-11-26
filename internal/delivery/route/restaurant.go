package route

import (
	protocomment "banners_oto/gen/comment"
	protorest "banners_oto/gen/rest"
	protouser "banners_oto/gen/user"
	"banners_oto/internal/delivery/metrics"
	repofood "banners_oto/internal/repository/food"
	reposearch "banners_oto/internal/repository/search"
	uscomment "banners_oto/internal/usecase/comment"
	usfood "banners_oto/internal/usecase/food"
	rest "banners_oto/internal/usecase/restaurants"
	ussearch "banners_oto/internal/usecase/search"
	ususer "banners_oto/internal/usecase/user"
	"banners_oto/microservices"
	"banners_oto/services"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func AddRestRouter(mux *mux.Router, cluster *services.Cluster, clients *microservices.Clients, logger *zap.Logger, metrics *metrics.Metrics) {
	repoSearch := reposearch.NewRepoLayer(cluster.PsqlClient, metrics)
	repoFood := repofood.NewRepo(cluster.PsqlClient, metrics)
	usecaseFood := usfood.NewUsecaseLayer(repoFood)
	usecaseSearch := ussearch.NewUsecaseLayer(repoSearch)
	// init user grpc client
	grpcUser := protouser.NewUserManagerClient(clients.UserConn)
	usecaseUser := ususer.NewUsecaseLayer(grpcUser, metrics)

	//init rest grpc client
	grpcRest := protorest.NewRestWorkerClient(clients.RestConn)
	usecaseRest := rest.NewUsecaseLayer(grpcRest, metrics)

	// init comment grpc client
	grpcComment := protocomment.NewCommentWorkerClient(clients.CommentConn)
	usecaseComment := uscomment.NewUseCaseLayer(grpcComment, grpcUser, metrics)

	deliveryRest := dRest.NewRestaurantHandler(usecaseRest, usecaseFood, usecaseUser, logger)
	deliveryComment := dComment.NewDelivery(usecaseComment, logger)
	deliverySearch := dSearch.NewDelivery(usecaseSearch, logger)

	mux.HandleFunc("/api/v1/search", deliverySearch.Search).Methods("GET").Name("restaurants-list")
	mux.HandleFunc("/api/v1/restaurants", deliveryRest.RestaurantList).Methods("GET").Name("restaurants-list")
	mux.HandleFunc("/api/v1/restaurants/{id}", deliveryRest.RestaurantById).Methods("GET").Name("restaurants-detail")
	mux.HandleFunc("/api/v1/restaurants/{id}/comment", deliveryComment.CreateComment).Methods("POST").Name("create-comment")
	mux.HandleFunc("/api/v1/restaurants/{id}/comment/{com_id}", deliveryComment.DeleteComment).Methods("DELETE").Name("delete-comment")
	mux.HandleFunc("/api/v1/restaurants/{id}/comments", deliveryComment.GetComments).Methods("GET").Name("comments-list")
	mux.HandleFunc("/api/v1/category", deliveryRest.CategoryList).Methods("GET").Name("category-list")
	mux.HandleFunc("/api/v1/recomendation", deliveryRest.Recomendation).Methods("GET").Name("recomendation")
}
