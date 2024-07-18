package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func Init(uri string) error {
	opt, err := redis.ParseURL(uri)
	if err != nil {
		return err
	}
	client = redis.NewClient(opt)
	return nil
}

func Set(key string, value string, expiration time.Duration) error {
	return client.Set(context.Background(), key, value, expiration).Err()
}

func Get(key string) (string, error) {
	return client.Get(context.Background(), key).Result()
}
