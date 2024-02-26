package storage

import (
	"context"

	"github.com/minio/minio-go/v7"
)

type Storage interface {
	UploadFile(ctx context.Context, localPath, destinationPath, bucket string) error
	GetFile(ctx context.Context, filename, bucket string) error
}

type storage struct {
	client *minio.Client
}

func New(minioClient *minio.Client) *storage {
	return &storage{client: minioClient}
}

func (s *storage) UploadFile(ctx context.Context, localPath, destinationPath, bucket string) error {
	return nil
}

func (s *storage) GetFile(ctx context.Context, filename, bucket string) error {
	return nil
}
