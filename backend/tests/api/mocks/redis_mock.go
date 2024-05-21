package mocks

import (
	"time"

	"github.com/GenerateNU/sac/backend/database/store"
)

// RedisMockClient implements StoreClientInterface
type RedisMockClient struct{}

func NewRedisMockClient() *RedisMockClient {
	return &RedisMockClient{}
}

func (r *RedisMockClient) Set(key string, value interface{}, expiration time.Duration) error {
	return nil
}

func (r *RedisMockClient) Get(key string) (string, error) {
	return "", nil
}

func (r *RedisMockClient) Del(key string) error {
	return nil
}

func (r *RedisMockClient) Exists(key string) (bool, error) {
	return false, nil
}

func (r *RedisMockClient) SAdd(key string, members ...interface{}) error {
	return nil
}

func (r *RedisMockClient) SIsMember(key string, member interface{}) (bool, error) {
	return false, nil
}

func (r *RedisMockClient) FlushAll() error {
	return nil
}

func (r *RedisMockClient) Close() error {
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

func (a *ActiveTokenMock) StoreRefreshToken(token string, userID string, expiry time.Duration) error {
	return nil
}

func (a *ActiveTokenMock) GetRefreshToken(token string) (string, error) {
	return "", nil
}

func (a *ActiveTokenMock) IsActive(token string) (bool, error) {
	return false, nil
}

func (a *ActiveTokenMock) DeleteRefreshToken(token string) error {
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

func (b *BlacklistMock) BlacklistToken(token string, expiry time.Duration) error {
	return nil
}

func (b *BlacklistMock) IsTokenBlacklisted(token string) (bool, error) {
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
