package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	client *redis.Client
}

func NewClient(uri string) (client *Client, err error) {
	opt, err := redis.ParseURL(uri)
	if err != nil {
		return nil, err
	}
	return &Client{client: redis.NewClient(opt)}, nil
}

func (redis *Client) Set(key string, value string, expiration time.Duration) error {
	return redis.client.Set(context.Background(), key, value, expiration).Err()
}

func (redis *Client) Get(key string) (string, error) {
	return redis.client.Get(context.Background(), key).Result()
}
