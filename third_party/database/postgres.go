package database

import (
	"context"
	"fmt"
	"net/url"

	"github.com/go-errors/errors"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Config used to init database connection
type Config struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
	SSLMode  string
}

// NewPostgres init database pool connection
func New(ctx context.Context, cfg Config) (*pgxpool.Pool, error) {
	// prepare connection url
	dbURL := fmt.Sprintf("postgres://%[1]s:%[2]s@%[3]s:%[4]s/%[5]s?sslmode=%[6]s",
		cfg.Username, url.QueryEscape(cfg.Password), cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)
	// init database connection
	dbPool, err := pgxpool.Connect(ctx, dbURL)
	if err != nil {
		return nil, errors.New(err)
	}

	// check  database connection
	err = dbPool.Ping(ctx)
	if err != nil {
		return nil, errors.New(err)
	}
	return dbPool, nil
}
