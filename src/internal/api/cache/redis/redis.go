package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(uri string) (client *Redis, err error) {
	opt, err := redis.ParseURL(uri)
	if err != nil {
		return nil, err
	}
	return &Redis{client: redis.NewClient(opt)}, nil
}

func (redis *Redis) Set(key string, value string, expiration time.Duration) error {
	return redis.client.Set(context.Background(), key, value, expiration).Err()
}

func (redis *Redis) Get(key string) (string, error) {
	return redis.client.Get(context.Background(), key).Result()
}

func (redis *Redis) Close() error {
	return redis.client.Close()
}
