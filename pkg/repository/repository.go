package repository

import (
	"context"

	"github.com/atadzan/simple-crud/pkg/models"
	"github.com/atadzan/simple-crud/pkg/repository/cache"
	"github.com/atadzan/simple-crud/pkg/repository/db"
	"github.com/atadzan/simple-crud/pkg/repository/storage"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
)

type Repo interface {
	GetAll(ctx context.Context, page, limit uint32) ([]models.BookList, error)
	GetById(ctx context.Context, id uint32) (models.BookDetailed, error)
	Create(ctx context.Context, params models.CreateBookParams) error
	Update(ctx context.Context, params models.UpdateBookParams) error
	Delete(ctx context.Context, id uint32) error
}

type repo struct {
	db      db.Book
	cache   cache.Book
	storage storage.Storage
}

func New(dbClient *pgxpool.Pool, minioClient *minio.Client, cacheClient *redis.Client) *repo {
	return &repo{
		db:      db.New(dbClient),
		storage: storage.New(minioClient),
		cache:   cache.New(cacheClient),
	}
}

func (r *repo) GetAll(ctx context.Context, page, limit uint32) ([]models.BookList, error) {
	return nil, nil
}
func (r *repo) GetById(ctx context.Context, id uint32) (models.BookDetailed, error) {
	return models.BookDetailed{}, nil
}
func (r *repo) Create(ctx context.Context, params models.CreateBookParams) error {
	return nil
}
func (r *repo) Update(ctx context.Context, params models.UpdateBookParams) error {
	return nil
}
func (r *repo) Delete(ctx context.Context, id uint32) error {
	return nil
}
