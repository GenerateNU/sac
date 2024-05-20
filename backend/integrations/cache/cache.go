package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheClientInterface interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (string, error)
	Del(key string) error
	Exists(key string) (bool, error)
	SAdd(key string, members ...interface{}) error
	SIsMember(key string, member interface{}) (bool, error)
}

type RedisClient struct {
	Client *redis.Client
	Ctx    context.Context
}

func NewRedisClient(addr, password string, db int) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisClient{
		Client: client,
		Ctx:    context.Background(),
	}
}

func (r *RedisClient) Set(key string, value interface{}, expiration time.Duration) error {
	return r.Client.Set(r.Ctx, key, value, expiration).Err()
}

func (r *RedisClient) Get(key string) (string, error) {
	return r.Client.Get(r.Ctx, key).Result()
}

func (r *RedisClient) Del(key string) error {
	return r.Client.Del(r.Ctx, key).Err()
}

func (r *RedisClient) Exists(key string) (bool, error) {
	result, err := r.Client.Exists(r.Ctx, key).Result()
	return result > 0, err
}

func (r *RedisClient) SAdd(key string, members ...interface{}) error {
	return r.Client.SAdd(r.Ctx, key, members...).Err()
}

func (r *RedisClient) SIsMember(key string, member interface{}) (bool, error) {
	return r.Client.SIsMember(r.Ctx, key, member).Result()
}
