package repository

import (
	"context"
	"log"
	"mime/multipart"

	"github.com/atadzan/simple-crud/pkg/models"
	"github.com/atadzan/simple-crud/pkg/repository/cache"
	"github.com/atadzan/simple-crud/pkg/repository/db"
	"github.com/atadzan/simple-crud/pkg/repository/storage"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
)

type Repo interface {
	// authorization methods
	Register(ctx context.Context, params models.AuthParams) error
	GetAuthorId(ctx context.Context, params models.AuthParams) (uint32, error)

	// Book methods
	GetGenres(ctx context.Context) ([]models.Genre, error)
	GetAll(ctx context.Context, params models.BooksParams) ([]models.BookList, error)
	GetById(ctx context.Context, id int) (models.BookDetailed, error)
	Create(ctx context.Context, params models.CreateBookParams) error
	Update(ctx context.Context, params models.UpdateBookParams) error
	Search(ct context.Context, params models.SearchParams) ([]models.BookList, error)
	Delete(ctx context.Context, id, authorId uint32) error

	UploadFile(ctx context.Context, header *multipart.FileHeader) error
	GetFile(ctx context.Context, filename string) (models.FileResponse, error)
}

type repo struct {
	db      db.DB
	cache   cache.Book
	storage storage.Storage
}

func New(dbClient *pgxpool.Pool, minioClient *minio.Client, cacheClient *redis.Client, domain string) *repo {
	return &repo{
		db:      db.New(dbClient, domain),
		storage: storage.New(minioClient),
		cache:   cache.New(cacheClient),
	}
}

func (r *repo) Register(ctx context.Context, params models.AuthParams) error {
	return r.db.Register(ctx, params)
}

func (r *repo) GetAuthorId(ctx context.Context, params models.AuthParams) (uint32, error) {
	return r.db.GetAuthorId(ctx, params)
}

func (r *repo) GetGenres(ctx context.Context) (genres []models.Genre, err error) {
	genres = make([]models.Genre, 0)

	// First get genres from cache
	genres, err = r.cache.GetGenres(ctx)

	// if it occurs error or not found go to database
	if err != nil || len(genres) == 0 {
		genres, err = r.db.GetGenres(ctx)
		if err != nil {
			return
		}

		// after successfully fetch from db set to cache
		if err = r.cache.SaveGenres(ctx, genres); err != nil {
			log.Println(err)
		}
	}
	return
}

func (r *repo) GetAll(ctx context.Context, params models.BooksParams) ([]models.BookList, error) {
	return r.db.GetAll(ctx, params)
}

func (r *repo) Search(ctx context.Context, params models.SearchParams) ([]models.BookList, error) {
	return r.db.Search(ctx, params)
}

func (r *repo) GetById(ctx context.Context, id int) (book models.BookDetailed, err error) {
	book, err = r.cache.GetById(ctx, id)
	if err != nil || len(book.Title) == 0 {
		book, err = r.db.GetById(ctx, id)
		if err != nil {
			return
		}
		// TODO Fix problem
		//if err = r.cache.Set(ctx, book); err != nil {
		//	log.Println(err)
		//}
	}

	return
}

func (r *repo) Create(ctx context.Context, params models.CreateBookParams) error {
	return r.db.Create(ctx, params)
}

func (r *repo) Update(ctx context.Context, params models.UpdateBookParams) error {
	return r.db.Update(ctx, params)
}

func (r *repo) Delete(ctx context.Context, id, authorId uint32) error {
	if err := r.db.Delete(ctx, id, authorId); err != nil {
		return err
	}
	if err := r.cache.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

func (r *repo) GetFile(ctx context.Context, filename string) (models.FileResponse, error) {
	return r.storage.GetFile(ctx, filename)
}

func (r *repo) UploadFile(ctx context.Context, header *multipart.FileHeader) error {
	return r.storage.UploadFile(ctx, header)
}
