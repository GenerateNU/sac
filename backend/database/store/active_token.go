package store

import (
	"context"
	"time"
)

type ActiveTokenInterface interface {
	StoreRefreshToken(token string, userID string, expiry time.Duration) error
	IsActive(token string) (bool, error)
	GetRefreshToken(token string) (string, error)
	DeleteRefreshToken(token string) error
}

type ActiveToken struct {
	StoreClient StoreClientInterface
	ctx         context.Context
}

func NewActiveToken(storeClient StoreClientInterface, ctx context.Context) *ActiveToken {
	return &ActiveToken{
		StoreClient: storeClient,
		ctx:         ctx,
	}
}

func (a *ActiveToken) StoreRefreshToken(token string, userID string, expiry time.Duration) error {
	return a.StoreClient.Set(a.ctx, token, userID, expiry)
}

func (a *ActiveToken) GetRefreshToken(token string) (string, error) {
	return a.StoreClient.Get(a.ctx, token)
}

func (a *ActiveToken) IsActive(token string) (bool, error) {
	return a.StoreClient.Exists(a.ctx, token)
}

func (a *ActiveToken) DeleteRefreshToken(token string) error {
	return a.StoreClient.Del(a.ctx, token)
}
