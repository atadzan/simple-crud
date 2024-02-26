package cache

import (
	"context"

	"github.com/go-errors/errors"
	"github.com/redis/go-redis/v9"
)

type Params struct {
	Host     string
	Port     string
	Password string
	DB       int
}

// NewRedisClient - establishing cache client connection with cache server
func NewRedisClient(ctx context.Context, cfg Params) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, errors.New(err)
	}
	return client, nil
}
