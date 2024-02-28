package db

import (
	"context"

	"github.com/atadzan/simple-crud/pkg/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Book interface {
	GetAll(ctx context.Context, page, limit uint32) ([]models.BookList, error)
	GetById(ctx context.Context, id uint32) (models.BookDetailed, error)
	Create(ctx context.Context, params models.CreateBookParams) error
	Update(ctx context.Context, params models.UpdateBookParams) error
	Delete(ctx context.Context, id uint32) error
}

type db struct {
	dbClient *pgxpool.Pool
}

func New(dbClient *pgxpool.Pool) *db {
	return &db{dbClient: dbClient}
}

func (d *db) GetAll(ctx context.Context, page, limit uint32) ([]models.BookList, error) {
	return nil, nil
}

func (d *db) GetById(ctx context.Context, id uint32) (models.BookDetailed, error) {
	return models.BookDetailed{}, nil
}

func (d *db) Create(ctx context.Context, params models.CreateBookParams) error {
	return nil
}

func (d *db) Update(ctx context.Context, params models.UpdateBookParams) error {
	return nil
}

func (d *db) Delete(ctx context.Context, id uint32) error {
	return nil
}
