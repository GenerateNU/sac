package store

import (
	"context"
	"time"
)

// LimiterInterface is an implementation of https://github.com/gofiber/storage
type LimiterInterface interface {
	Get(key string) ([]byte, error)
	Set(key string, val []byte, exp time.Duration) error
	Delete(key string) error
	Reset() error
	Close() error
}

type Limiter struct {
	StoreClient StoreClientInterface
}

func NewLimiter(storeClient StoreClientInterface) *Limiter {
	return &Limiter{
		StoreClient: storeClient,
	}
}

func (l *Limiter) Get(key string) ([]byte, error) {
	value, err := l.StoreClient.Get(context.Background(), key)
	if err != nil {
		return nil, err
	}

	return []byte(value), nil
}

func (l *Limiter) Set(key string, val []byte, exp time.Duration) error {
	return l.StoreClient.Set(context.Background(), key, string(val), exp)
}

func (l *Limiter) Delete(key string) error {
	return l.StoreClient.Del(context.Background(), key)
}

func (l *Limiter) Reset() error {
	return l.StoreClient.FlushAll(context.Background())
}

func (l *Limiter) Close() error {
	return l.StoreClient.Close(context.Background())
}
