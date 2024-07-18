package cache

import (
	"github.com/quinntas/go-gin-template/src/internal/api/cache/redis"
	"time"
)

func Init(uri string) error {
	err := redis.Init(uri)
	if err != nil {
		return err
	}
	return nil
}

func Set(key string, value string, expiration time.Duration) error {
	return redis.Set(key, value, expiration)
}

func Get(key string) (string, error) {
	return redis.Get(key)
}
