package cache

import "github.com/redis/go-redis/v9"

type Book interface {
	Get()
}

type cache struct {
	client *redis.Client
}

func New(cacheClient *redis.Client) *cache {
	return &cache{client: cacheClient}
}

func (c *cache) Get() {
	return
}
