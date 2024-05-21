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
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (string, error)
	Del(key string) error
	Exists(key string) (bool, error)
	SAdd(key string, members ...interface{}) error
	SIsMember(key string, member interface{}) (bool, error)
	FlushAll() error
	Close() error
}

type RedisClient struct {
	Client *redis.Client
	Ctx    context.Context
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

func (r *RedisClient) FlushAll() error {
	return r.Client.FlushAll(r.Ctx).Err()
}

func (r *RedisClient) Close() error {
	return r.Client.Close()
}
