package storage

import (
	"context"
	"mime/multipart"

	"github.com/atadzan/simple-crud/pkg/models"
	"github.com/minio/minio-go/v7"
)

type Storage interface {
	UploadFile(ctx context.Context, header *multipart.FileHeader) error
	GetFile(ctx context.Context, filename string) (models.FileResponse, error)
}

type storage struct {
	client *minio.Client
}

func New(minioClient *minio.Client) *storage {
	return &storage{client: minioClient}
}

func (s *storage) UploadFile(ctx context.Context, header *multipart.FileHeader) error {
	file, err := header.Open()
	if err != nil {
		return err
	}
	// First we check bucket if bucket doesn't exist, we create new one
	isExists, err := s.client.BucketExists(ctx, "images")
	if err != nil || !isExists {
		if err = s.client.MakeBucket(ctx, "images", minio.MakeBucketOptions{}); err != nil {
			return err
		}
	}
	// Upload file to storage
	_, err = s.client.PutObject(ctx, "images", header.Filename, file, -1, minio.PutObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (s *storage) GetFile(ctx context.Context, filename string) (models.FileResponse, error) {
	file, err := s.client.GetObject(ctx, "images", filename, minio.GetObjectOptions{})
	if err != nil {
		return models.FileResponse{}, err
	}
	stat, err := file.Stat()

	if err != nil {
		return models.FileResponse{}, err
	}
	return models.FileResponse{
		Reader: file,
		Size:   stat.Size,
	}, nil
}
