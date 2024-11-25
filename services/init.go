package services

import (
	"database/sql"

	"github.com/minio/minio-go"
	"go.uber.org/zap"
)

type Cluster struct {
	PsqlClient  *sql.DB
	MinioClient *minio.Client
}

func Init(logger *zap.Logger) *Cluster {
	return &Cluster{
		PsqlClient:  postgres.Init(logger),
		MinioClient: minios3.Init(logger),
	}
}
