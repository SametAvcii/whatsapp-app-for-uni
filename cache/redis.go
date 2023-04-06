package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v9"
)

type ICache interface {
	Get(ctx context.Context, key string) bool
	Set(ctx context.Context, key string, list interface{}, ex int64) error
}

type Cache struct {
	client *redis.Client
}

func NewCache(client *redis.Client) ICache {
	return &Cache{client: client}
}

func (c *Cache) Get(ctx context.Context, key string) bool {
	err := c.client.Get(ctx, key).Err()
	if err != nil {
		return false
	}
	return true

}

func (c *Cache) Set(ctx context.Context, key string, list interface{}, ex int64) error {
	jsondata, err := json.Marshal(list)
	if err != nil {
		return err
	}
	err = c.client.Set(ctx, key, jsondata, time.Duration(ex*int64(time.Second))).Err()
	if err != nil {
		return err
	}
	return nil
}
