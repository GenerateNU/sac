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
	ctx         context.Context
}

func NewLimiter(storeClient StoreClientInterface, ctx context.Context) *Limiter {
	return &Limiter{
		StoreClient: storeClient,
		ctx:         ctx,
	}
}

func (l *Limiter) Get(key string) ([]byte, error) {
	value, err := l.StoreClient.Get(l.ctx, key)
	if err != nil {
		return nil, err
	}

	return []byte(value), nil
}

func (l *Limiter) Set(key string, val []byte, exp time.Duration) error {
	return l.StoreClient.Set(l.ctx, key, string(val), exp)
}

func (l *Limiter) Delete(key string) error {
	return l.StoreClient.Del(l.ctx, key)
}

func (l *Limiter) Reset() error {
	return l.StoreClient.FlushAll(l.ctx)
}

func (l *Limiter) Close() error {
	return l.StoreClient.Close()
}
