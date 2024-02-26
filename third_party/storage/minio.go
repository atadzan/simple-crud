package storage

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-errors/errors"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Params struct {
	StorageId         string
	Endpoint          string
	AccessKeyId       string
	SecretAccessKeyId string
}

func checkStorageHealth(params Params) error {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(fmt.Sprintf("http://%s/minio/health/live", params.Endpoint))
	if err != nil || resp.StatusCode != http.StatusOK {
		return errors.New("Cannot connect to storage")
	}
	return nil
}

func New(cfg Params) (*minio.Client, error) {
	minioClient, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyId, cfg.SecretAccessKeyId, ""),
		Secure: false,
		Region: "us-east-1",
	})
	if err != nil {
		return nil, errors.New(fmt.Errorf("failed to create minio client. err: %w", err))
	}
	if err = checkStorageHealth(cfg); err != nil {
		return nil, err
	}

	return minioClient, nil
}
