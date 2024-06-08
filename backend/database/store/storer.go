package store

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/constants"
	"github.com/redis/go-redis/v9"
)

// an implementation of https://github.com/gofiber/storage
type Storer interface {
	Get(key string) ([]byte, error)
	Set(key string, val []byte, exp time.Duration) error
	Delete(key string) error
	Reset() error
	Close() error
}

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(settings config.RedisSettings) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Username:              settings.Username,
		Password:              settings.Password.Expose(),
		Addr:                  fmt.Sprintf("%s:%d", settings.Host, settings.Port),
		DB:                    settings.DB,
		PoolSize:              10 * runtime.GOMAXPROCS(0),
		MaxActiveConns:        constants.REDIS_MAX_OPEN_CONNECTIONS,
		MaxIdleConns:          constants.REDIS_MAX_IDLE_CONNECTIONS,
		ContextTimeoutEnabled: true,
	})

	return &RedisClient{
		client: client,
	}
}

func (r *RedisClient) Get(key string) ([]byte, error) {
	return r.client.Get(context.Background(), key).Bytes()
}

func (r *RedisClient) Set(key string, val []byte, exp time.Duration) error {
	return r.client.Set(context.Background(), key, val, exp).Err()
}

func (r *RedisClient) Delete(key string) error {
	return r.client.Del(context.Background(), key).Err()
}

func (r *RedisClient) Reset() error {
	return r.client.FlushAll(context.Background()).Err()
}

func (r *RedisClient) Close() error {
	return r.client.Close()
}
