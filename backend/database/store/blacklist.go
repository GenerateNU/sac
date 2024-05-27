package store

import (
	"context"
	"time"

	"github.com/GenerateNU/sac/backend/constants"
)

type BlacklistInterface interface {
	BlacklistToken(token string, expiry time.Duration) error
	IsTokenBlacklisted(token string) (bool, error)
}

type Blacklist struct {
	StoreClient StoreClientInterface
	ctx         context.Context
}

func NewBlacklist(storeClient StoreClientInterface, ctx context.Context) *Blacklist {
	return &Blacklist{
		StoreClient: storeClient,
		ctx:         ctx,
	}
}

func (b *Blacklist) BlacklistToken(token string, expiry time.Duration) error {
	err := b.StoreClient.SAdd(b.ctx, constants.TOKEN_BLACKLIST_SKEY, token)
	if err != nil {
		return err
	}

	return b.StoreClient.Set(b.ctx, token, constants.TOKEN_BLACKLIST_KEY, expiry)
}

func (b *Blacklist) IsTokenBlacklisted(token string) (bool, error) {
	return b.StoreClient.SIsMember(b.ctx, constants.TOKEN_BLACKLIST_SKEY, token)
}
