package cache

import (
	"github.com/atadzan/simple-crud/pkg/models"
	"github.com/redis/go-redis/v9"
)

type Book interface {
	Get(id uint32) (models.BookDetailed, error)
}

type cache struct {
	client *redis.Client
}

func New(cacheClient *redis.Client) *cache {
	return &cache{client: cacheClient}
}

func (c *cache) Get(id uint32) (models.BookDetailed, error) {

	return models.BookDetailed{}, nil
}
