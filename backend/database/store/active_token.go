package store

import "time"

type ActiveTokenInterface interface {
	StoreRefreshToken(token string, userID string, expiry time.Duration) error
	IsActive(token string) (bool, error)
	GetRefreshToken(token string) (string, error)
	DeleteRefreshToken(token string) error
}

type ActiveToken struct {
	StoreClient StoreClientInterface
}

func NewActiveToken(storeClient StoreClientInterface) *ActiveToken {
	return &ActiveToken{
		StoreClient: storeClient,
	}
}

func (a *ActiveToken) StoreRefreshToken(token string, userID string, expiry time.Duration) error {
	return a.StoreClient.Set(token, userID, expiry)
}

func (a *ActiveToken) GetRefreshToken(token string) (string, error) {
	return a.StoreClient.Get(token)
}

func (a *ActiveToken) IsActive(token string) (bool, error) {
	return a.StoreClient.Exists(token)
}

func (a *ActiveToken) DeleteRefreshToken(token string) error {
	return a.StoreClient.Del(token)
}
