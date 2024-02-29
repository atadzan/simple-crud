package cache

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/atadzan/simple-crud/pkg/models"
	"github.com/redis/go-redis/v9"
)

type Book interface {
	GetGenres(ctx context.Context) ([]models.Genre, error)
	SaveGenres(ctx context.Context, genres []models.Genre) error

	GetById(ctx context.Context, id int) (models.BookDetailed, error)
	Set(ctx context.Context, book models.BookDetailed) error
	Delete(ctx context.Context, id uint32) error
}

type cache struct {
	client *redis.Client
}

func New(cacheClient *redis.Client) *cache {
	return &cache{client: cacheClient}
}

func (c *cache) GetGenres(ctx context.Context) (genres []models.Genre, err error) {
	resp, err := c.client.Get(ctx, "genres").Bytes()
	if err != nil {
		log.Println(err)
		return
	}
	if err = json.Unmarshal(resp, &genres); err != nil {
		log.Println(err)
		return
	}

	return
}

func (c *cache) SaveGenres(ctx context.Context, genres []models.Genre) (err error) {
	b, err := json.Marshal(&genres)
	if err != nil {
		return
	}
	if err = c.client.Set(ctx, "genres", b, 15*time.Minute).Err(); err != nil {
		return err
	}
	return
}

func (c *cache) GetById(ctx context.Context, id int) (book models.BookDetailed, err error) {
	if err = c.client.HGetAll(ctx, string(id)).Scan(&book); err != nil {
		log.Println(err)
		return
	}
	return
}

func (c *cache) Set(ctx context.Context, book models.BookDetailed) error {
	if err := c.client.HSet(ctx, string(book.ID), book).Err(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (c *cache) Delete(ctx context.Context, id uint32) error {
	if err := c.client.Del(ctx, string(id)).Err(); err != nil {
		return err
	}
	return nil
}
