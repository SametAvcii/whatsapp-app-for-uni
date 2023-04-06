package redis

import (
	"context"

	"github.com/go-redis/redis/v9"
)

var ctx = context.Background()

func NewClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return rdb
}
