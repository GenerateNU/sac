package store

import "time"

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

func NewLimiter(storeCleint StoreClientInterface) *Limiter {
	return &Limiter{
		StoreClient: storeCleint,
	}
}

func (l *Limiter) Get(key string) ([]byte, error) {
	value, err := l.StoreClient.Get(key)
	if err != nil {
		return nil, err
	}

	return []byte(value), nil
}

func (l *Limiter) Set(key string, val []byte, exp time.Duration) error {
	return l.StoreClient.Set(key, val, exp)
}

func (l *Limiter) Delete(key string) error {
	return l.StoreClient.Del(key)
}

func (l *Limiter) Reset() error {
	return l.StoreClient.FlushAll()
}

func (l *Limiter) Close() error {
	return l.StoreClient.Close()
}
