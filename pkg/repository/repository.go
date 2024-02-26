package repository

import (
	"context"

	"github.com/atadzan/simple-crud/pkg/repository/cache"
	"github.com/atadzan/simple-crud/pkg/repository/db"
	"github.com/atadzan/simple-crud/pkg/repository/storage"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
)

type Repo interface {
	GetAll(ctx context.Context, page, limit uint32) (models.BookListItem, error)
	GetById(ctx context.Context, uniqueId string) (models.BookDetailedInfo, error)
	Create(ctx context.Context, params models.CreateParams) error
	Update(ctx context.Context, params models.UpdateParams) error
	Delete(ctx context.Context, uniqueId string) error
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

func (r *repo) GetAll(ctx context.Context, page, limit uint32) (models.BookListItem, error) {
	return
}
func (r *repo) GetById(ctx context.Context, uniqueId string) (models.BookDetailedInfo, error) {
	return
}
func (r *repo) Create(ctx context.Context, params models.CreateParams) error {
	return
}
func (r *repo) Update(ctx context.Context, params models.UpdateParams) error {
	return
}
func (r *repo) Delete(ctx context.Context, uniqueId string) error {
	return
}
