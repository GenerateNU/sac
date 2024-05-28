package store

import (
	"context"
	"time"
)

type ActiveTokenInterface interface {
	StoreRefreshToken(ctx context.Context, token string, userID string, expiry time.Duration) error
	IsActive(ctx context.Context, token string) (bool, error)
	GetRefreshToken(ctx context.Context, token string) (string, error)
	DeleteRefreshToken(ctx context.Context, token string) error
}

type ActiveToken struct {
	StoreClient StoreClientInterface
}

func NewActiveToken(storeClient StoreClientInterface) *ActiveToken {
	return &ActiveToken{
		StoreClient: storeClient,
	}
}

func (a *ActiveToken) StoreRefreshToken(ctx context.Context, token string, userID string, expiry time.Duration) error {
	return a.StoreClient.Set(ctx, token, userID, expiry)
}

func (a *ActiveToken) GetRefreshToken(ctx context.Context, token string) (string, error) {
	return a.StoreClient.Get(ctx, token)
}

func (a *ActiveToken) IsActive(ctx context.Context, token string) (bool, error) {
	return a.StoreClient.Exists(ctx, token)
}

func (a *ActiveToken) DeleteRefreshToken(ctx context.Context, token string) error {
	return a.StoreClient.Del(ctx, token)
}
