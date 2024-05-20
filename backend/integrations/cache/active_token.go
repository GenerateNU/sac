package cache

import "time"

type ActiveTokenInterface interface {
	StoreRefreshToken(token string, userID string, expiry time.Duration) error
}

type ActiveToken struct {
	CacheClient CacheClientInterface
}

func NewActiveToken(cacheClient CacheClientInterface) *ActiveToken {
	return &ActiveToken{
		CacheClient: cacheClient,
	}
}

func (a *ActiveToken) StoreRefreshToken(token string, userID string, expiry time.Duration) error {
	return a.CacheClient.Set(token, userID, expiry)
}
