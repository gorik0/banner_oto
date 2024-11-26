package minios3

import (
	cnst "banners_oto/internal/utils/constants"
	"bytes"
	"context"

	"github.com/minio/minio-go/v7"
)

type Repo interface {
	UploadImageByEmail(ctx context.Context, buffer *bytes.Buffer, filename string, filesize int64, mimeType string) error
}

type RepoLayer struct {
	minio *minio.Client
}

func NewRepoLayer(minioClient *minio.Client) Repo {
	return &RepoLayer{
		minio: minioClient,
	}
}

func (repo *RepoLayer) UploadImageByEmail(ctx context.Context, buffer *bytes.Buffer, filename string, filesize int64, mimeType string) error {
	_, err := repo.minio.PutObject(ctx, cnst.BucketUser, filename, buffer, filesize, minio.PutObjectOptions{ContentType: mimeType})
	return err
}
