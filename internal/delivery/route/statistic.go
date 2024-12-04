package route

import (
	sessionproto "banners_oto/gen/session"
	protouser "banners_oto/gen/user"
	"banners_oto/internal/delivery/metrics"
	"banners_oto/internal/delivery/statistic"
	repostatistic "banners_oto/internal/repository/statistic"
	ussession "banners_oto/internal/usecase/session"
	usstatistic "banners_oto/internal/usecase/statistic"
	ususer "banners_oto/internal/usecase/user"
	"banners_oto/microservices"
	"banners_oto/services"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func AddQuizRouter(mux *mux.Router, cluster *services.Cluster, clients *microservices.Clients, logger *zap.Logger, metrics *metrics.Metrics) {
	repoQuiz := repostatistic.NewRepoLayer(cluster.PsqlClient, metrics)
	// init grpc client interface
	grpcSessionClient := sessionproto.NewSessionManagerClient(clients.SessionConn)
	usecaseSession := ussession.NewUsecaseLayer(grpcSessionClient, metrics)
	// init grpc user interface
	grpcUserClient := protouser.NewUserManagerClient(clients.UserConn)
	usecaseUser := ususer.NewUsecaseLayer(grpcUserClient, metrics)

	usecaseQuiz := usstatistic.NewUsecaseLayer(repoQuiz)

	deliveryQuiz := statistic.NewDeliveryLayer(usecaseQuiz, usecaseUser, usecaseSession, logger, metrics)

	mux.HandleFunc("/api/v1/quiz/stats", deliveryQuiz.GetStatistic).Methods("GET").Name("quiz-stats")
	mux.HandleFunc("/api/v1/quiz/questions", deliveryQuiz.GetQuestions).Methods("GET").Name("get-questions")
	mux.HandleFunc("/api/v1/quiz/question/rating", deliveryQuiz.AddAnswer).Methods("POST").Name("add-question-rating")
}
