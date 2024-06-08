package mocks

import (
	"time"
)

type RedisClient struct{}

func NewRedisClient() *RedisClient {
	return &RedisClient{}
}

func (r *RedisClient) Get(key string) ([]byte, error) {
	return []byte{}, nil
}

func (r *RedisClient) Set(key string, val []byte, exp time.Duration) error {
	return nil
}

func (r *RedisClient) Delete(key string) error {
	return nil
}

func (r *RedisClient) Reset() error {
	return nil
}

func (r *RedisClient) Close() error {
	return nil
}
