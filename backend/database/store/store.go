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
	SAdd(ctx context.Context, key string, members ...interface{}) error
	SIsMember(ctx context.Context, key string, member interface{}) (bool, error)
	FlushAll(ctx context.Context) error
	Close() error
}

type RedisClient struct {
	Client *redis.Client
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
		Client: client,
	}
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.Client.Set(ctx, key, value, expiration).Err()
}

func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

func (r *RedisClient) Del(ctx context.Context, key string) error {
	return r.Client.Del(ctx, key).Err()
}

func (r *RedisClient) Exists(ctx context.Context, key string) (bool, error) {
	result, err := r.Client.Exists(ctx, key).Result()
	return result > 0, err
}

func (r *RedisClient) SAdd(ctx context.Context, key string, members ...interface{}) error {
	return r.Client.SAdd(ctx, key, members...).Err()
}

func (r *RedisClient) SIsMember(ctx context.Context, key string, member interface{}) (bool, error) {
	return r.Client.SIsMember(ctx, key, member).Result()
}

func (r *RedisClient) FlushAll(ctx context.Context) error {
	return r.Client.FlushAll(ctx).Err()
}

func (r *RedisClient) Close() error {
	return r.Client.Close()
}
