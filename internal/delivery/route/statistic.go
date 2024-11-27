package route

import (
	ucSession "2024_1_kayros/internal/usecase/session"
	sessionproto "banners_oto/gen/session"
	"banners_oto/internal/delivery/metrics"
	"banners_oto/microservices"
	"banners_oto/services"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func AddQuizRouter(mux *mux.Router, cluster *services.Cluster, clients *microservices.Clients, logger *zap.Logger, metrics *metrics.Metrics) {
	repoQuiz := rQuiz.NewRepoLayer(cluster.PsqlClient, metrics)
	// init grpc client interface
	grpcSessionClient := sessionproto.NewSessionManagerClient(clients.SessionConn)
	usecaseSession := ucSession.NewUsecaseLayer(grpcSessionClient, metrics)
	// init grpc user interface
	grpcUserClient := userproto.NewUserManagerClient(clients.UserConn)
	usecaseUser := uUser.NewUsecaseLayer(grpcUserClient, metrics)

	usecaseQuiz := uQuiz.NewUsecaseLayer(repoQuiz)

	deliveryQuiz := dQuiz.NewDeliveryLayer(usecaseQuiz, usecaseUser, usecaseSession, logger, metrics)

	mux.HandleFunc("/api/v1/quiz/stats", deliveryQuiz.GetStatistic).Methods("GET").Name("quiz-stats")
	mux.HandleFunc("/api/v1/quiz/questions", deliveryQuiz.GetQuestions).Methods("GET").Name("get-questions")
	mux.HandleFunc("/api/v1/quiz/question/rating", deliveryQuiz.AddAnswer).Methods("POST").Name("add-question-rating")
}
