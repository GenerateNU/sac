package mocks

import (
	"context"
	"time"

	"github.com/GenerateNU/sac/backend/database/store"
)

type RedisMockClient struct{}

func NewRedisMockClient() *RedisMockClient {
	return &RedisMockClient{}
}

func (r *RedisMockClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return nil
}

func (r *RedisMockClient) Get(ctx context.Context, key string) (string, error) {
	return "", nil
}

func (r *RedisMockClient) Delete(ctx context.Context, key string) error {
	return nil
}

func (r *RedisMockClient) Exists(ctx context.Context, key string) (bool, error) {
	return false, nil
}

func (r *RedisMockClient) SetAdd(ctx context.Context, key string, members ...interface{}) error {
	return nil
}

func (r *RedisMockClient) SetIsMember(ctx context.Context, key string, member interface{}) (bool, error) {
	return false, nil
}

func (r *RedisMockClient) FlushAll(ctx context.Context) error {
	return nil
}

func (r *RedisMockClient) Close(ctx context.Context) error {
	return nil
}

type LimiterMock struct {
	store.Storer
}

func NewLimiterMock(storer store.Storer) *LimiterMock {
	return &LimiterMock{
		storer,
	}
}

func (l *LimiterMock) Get(key string) ([]byte, error) {
	return nil, nil
}

func (l *LimiterMock) Set(key string, val []byte, exp time.Duration) error {
	return nil
}

func (l *LimiterMock) Delete(key string) error {
	return nil
}

func (l *LimiterMock) Reset() error {
	return nil
}

func (l *LimiterMock) Close() error {
	return nil
}

type RefresherMock struct {
	store.Storer
}

func NewRefresherMock(storer store.Storer) *RefresherMock {
	return &RefresherMock{
		storer,
	}
}

func (r *RefresherMock) Get(ctx context.Context, userID string) (string, error) {
	return "", nil
}

func (r *RefresherMock) Set(ctx context.Context, userID string, token string, expiry time.Duration) error {
	return nil
}

func (r *RefresherMock) Delete(ctx context.Context, userID string) error {
	return nil
}

func (r *RefresherMock) Exists(ctx context.Context, userID string) (bool, error) {
	return false, nil
}
