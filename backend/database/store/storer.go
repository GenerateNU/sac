package store

import (
	"context"
	"time"

	"github.com/GenerateNU/sac/backend/config"
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
	return &RedisClient{
		client: settings.Into(),
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
