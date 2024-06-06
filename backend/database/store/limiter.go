package store

import (
	"context"
	"time"
)

// Limiter is an implementation of https://github.com/gofiber/storage
type Limiter interface {
	Get(key string) ([]byte, error)
	Set(key string, val []byte, exp time.Duration) error
	Delete(key string) error
	Reset() error
	Close() error
}

type Limit struct {
	store Storer
}

func NewLimit(store Storer) *Limit {
	return &Limit{
		store: store,
	}
}

func (l *Limit) Get(key string) ([]byte, error) {
	value, err := l.store.Get(context.Background(), key)
	if err != nil {
		return nil, err
	}

	return []byte(value), nil
}

func (l *Limit) Set(key string, val []byte, exp time.Duration) error {
	return l.store.Set(context.Background(), key, string(val), exp)
}

func (l *Limit) Delete(key string) error {
	return l.store.Delete(context.Background(), key)
}

func (l *Limit) Reset() error {
	return l.store.FlushAll(context.Background())
}

func (l *Limit) Close() error {
	return l.store.Close(context.Background())
}
