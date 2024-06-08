package mocks

import (
	"context"
	"time"

	"github.com/GenerateNU/sac/backend/database/store"
)

type RedisClient struct{}

func NewRedisClient() *RedisClient {
	return &RedisClient{}
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return nil
}

func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	return "", nil
}

func (r *RedisClient) Delete(ctx context.Context, key string) error {
	return nil
}

func (r *RedisClient) Exists(ctx context.Context, key string) (bool, error) {
	return false, nil
}

func (r *RedisClient) SetAdd(ctx context.Context, key string, members ...interface{}) error {
	return nil
}

func (r *RedisClient) SetIsMember(ctx context.Context, key string, member interface{}) (bool, error) {
	return false, nil
}

func (r *RedisClient) FlushAll(ctx context.Context) error {
	return nil
}

func (r *RedisClient) Close(ctx context.Context) error {
	return nil
}

type Limiter struct {
	store.Storer
}

func NewLimiter(storer store.Storer) *Limiter {
	return &Limiter{
		storer,
	}
}

func (l *Limiter) Get(key string) ([]byte, error) {
	return nil, nil
}

func (l *Limiter) Set(key string, val []byte, exp time.Duration) error {
	return nil
}

func (l *Limiter) Delete(key string) error {
	return nil
}

func (l *Limiter) Reset() error {
	return nil
}

func (l *Limiter) Close() error {
	return nil
}

type Refresher struct {
	store.Storer
}

func NewRefresher(storer store.Storer) *Refresher {
	return &Refresher{
		storer,
	}
}

func (r *Refresher) Get(ctx context.Context, userID string) (string, error) {
	return "", nil
}

func (r *Refresher) Set(ctx context.Context, userID string, token string, expiry time.Duration) error {
	return nil
}

func (r *Refresher) Delete(ctx context.Context, userID string) error {
	return nil
}

func (r *Refresher) Exists(ctx context.Context, userID string) (bool, error) {
	return false, nil
}
