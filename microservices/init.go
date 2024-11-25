package microservices

import (
	"banners_oto/internal/delivery/metrics"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Clients struct {
	CommentConn *grpc.ClientConn
	RestConn    *grpc.ClientConn
	AuthConn    *grpc.ClientConn
	UserConn    *grpc.ClientConn
	SessionConn *grpc.ClientConn
}

func Init(logger *zap.Logger, m *metrics.Metrics) *Clients {

}
