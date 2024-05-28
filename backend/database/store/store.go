package store

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/GenerateNU/sac/backend/constants"
	m "github.com/garrettladley/mattress"
	"github.com/redis/go-redis/v9"
)

type StoreClientInterface interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
	Exists(ctx context.Context, key string) (bool, error)
	SetAdd(ctx context.Context, key string, members ...interface{}) error
	SetIsMember(ctx context.Context, key string, member interface{}) (bool, error)
	FlushAll(ctx context.Context) error
	Close(ctx context.Context) error
}

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(username, host string, port uint, password *m.Secret[string], db int) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Username:       username,
		Addr:           fmt.Sprintf("%s:%d", host, port),
		Password:       password.Expose(),
		DB:             db,
		PoolSize:       10 * runtime.GOMAXPROCS(0),
		MaxActiveConns: constants.REDIS_MAX_OPEN_CONNECTIONS,
		MaxIdleConns:   constants.REDIS_MAX_IDLE_CONNECTIONS,
	})

	return &RedisClient{
		client: client,
	}
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *RedisClient) Del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *RedisClient) Exists(ctx context.Context, key string) (bool, error) {
	result, err := r.client.Exists(ctx, key).Result()
	return result > 0, err
}

func (r *RedisClient) SetAdd(ctx context.Context, key string, members ...interface{}) error {
	return r.client.SAdd(ctx, key, members...).Err()
}

func (r *RedisClient) SetIsMember(ctx context.Context, key string, member interface{}) (bool, error) {
	return r.client.SIsMember(ctx, key, member).Result()
}

func (r *RedisClient) FlushAll(ctx context.Context) error {
	return r.client.FlushAll(ctx).Err()
}

func (r *RedisClient) Close(ctx context.Context) error {
	return r.client.Close()
}
