package db

import (
	"context"
	"fmt"

	"github.com/atadzan/simple-crud/pkg/models"
	"github.com/go-errors/errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	authorsTable = "authors"
	genresTable  = "genres"
	booksTable   = "books"
)

type Book interface {
	GetAll(ctx context.Context, page, limit uint32) ([]models.BookList, error)
	GetById(ctx context.Context, id uint32) (models.BookDetailed, error)
	Create(ctx context.Context, params models.CreateBookParams) error
	Update(ctx context.Context, params models.UpdateBookParams) error
	Delete(ctx context.Context, id uint32) error
}

type Auth interface {
	Register(ctx context.Context, params models.AuthParams) error
	GetAuthorId(ctx context.Context, params models.AuthParams) (uint32, error)
}

type DB interface {
	Book
	Auth
}
type db struct {
	dbClient *pgxpool.Pool
}

func New(dbClient *pgxpool.Pool) *db {
	return &db{dbClient: dbClient}
}

func (d *db) Register(ctx context.Context, params models.AuthParams) error {
	query := fmt.Sprintf(`INSERT INTO %s(username, password_hash) VALUES($1, $2)`, authorsTable)
	row, err := d.dbClient.Exec(ctx, query, params.Username, params.Password)
	if err != nil {
		return errors.New(err)
	}
	if row.RowsAffected() == 0 {
		return errors.New("operation failed")
	}
	return nil
}

func (d *db) GetAuthorId(ctx context.Context, params models.AuthParams) (authorID uint32, err error) {
	query := fmt.Sprintf(`SELECT id FROM %s WHERE username=$1 AND password_hash=$2`, authorsTable)
	if err = d.dbClient.QueryRow(ctx, query, params.Username, params.Password).Scan(&authorID); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = errors.New(ErrNotFound)
		} else {
			err = errors.New(err)
		}
		return
	}
	return
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
