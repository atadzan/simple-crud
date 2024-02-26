package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Book interface {
	GetAll(ctx context.Context, page, limit uint32) (models.BookListItem, error)
	GetById(ctx context.Context, uniqueId string) (models.BookDetailedInfo, error)
	Create(ctx context.Context, params models.CreateParams) error
	Update(ctx context.Context, params models.UpdateParams) error
	Delete(ctx context.Context, uniqueId string) error
}

type db struct {
	dbClient *pgxpool.Pool
}

func New(dbClient *pgxpool.Pool) *db {
	return &db{dbClient: dbClient}
}

func (d *db) GetAll(ctx context.Context, page, limit uint32) (models.BookListItem, error) {
	return nil, nil
}

func (d *db) GetById(ctx context.Context, uniqueId string) (models.BookDetailedInfo, error) {
	return nil, nil
}

func (d *db) Create(ctx context.Context, params models.CreateParams) error {
	return nil
}

func (d *db) Update(ctx context.Context, params models.UpdateParams) error {
	return nil
}

func (d *db) Delete(ctx context.Context, uniqueId string) error {
	return nil
}
