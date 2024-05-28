package mocks

import (
	"context"
	"time"

	"github.com/GenerateNU/sac/backend/database/store"
)

// RedisMockClient implements StoreClientInterface
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

func (r *RedisMockClient) Del(ctx context.Context, key string) error {
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

// ActiveTokenMock implements ActiveTokenInterface
type ActiveTokenMock struct {
	StoreClient store.StoreClientInterface
}

func NewActiveTokenMock(storeClient store.StoreClientInterface) *ActiveTokenMock {
	return &ActiveTokenMock{
		StoreClient: storeClient,
	}
}

func (a *ActiveTokenMock) StoreRefreshToken(ctx context.Context, token string, userID string, expiry time.Duration) error {
	return nil
}

func (a *ActiveTokenMock) GetRefreshToken(ctx context.Context, token string) (string, error) {
	return "", nil
}

func (a *ActiveTokenMock) IsActive(ctx context.Context, token string) (bool, error) {
	return false, nil
}

func (a *ActiveTokenMock) DeleteRefreshToken(ctx context.Context, token string) error {
	return nil
}

// BlacklistMock implements BlacklistInterface
type BlacklistMock struct {
	StoreClient store.StoreClientInterface
}

func NewBlacklistMock(storeClient store.StoreClientInterface) *BlacklistMock {
	return &BlacklistMock{
		StoreClient: storeClient,
	}
}

func (b *BlacklistMock) BlacklistToken(ctx context.Context, token string, expiry time.Duration) error {
	return nil
}

func (b *BlacklistMock) IsTokenBlacklisted(ctx context.Context, token string) (bool, error) {
	return false, nil
}

// LimiterMock implements LimiterInterface
type LimiterMock struct {
	StoreClient store.StoreClientInterface
}

func NewLimiterMock(storeClient store.StoreClientInterface) *LimiterMock {
	return &LimiterMock{
		StoreClient: storeClient,
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
