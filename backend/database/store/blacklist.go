package store

import (
	"context"
	"time"

	"github.com/GenerateNU/sac/backend/constants"
)

type BlacklistInterface interface {
	BlacklistToken(ctx context.Context, token string, expiry time.Duration) error
	IsTokenBlacklisted(ctx context.Context, token string) (bool, error)
}

type Blacklist struct {
	storeClient StoreClientInterface
}

func NewBlacklist(storeClient StoreClientInterface) *Blacklist {
	return &Blacklist{
		storeClient: storeClient,
	}
}

func (b *Blacklist) BlacklistToken(ctx context.Context, token string, expiry time.Duration) error {
	err := b.storeClient.SetAdd(ctx, constants.TOKEN_BLACKLIST_SKEY, token)
	if err != nil {
		return err
	}

	return b.storeClient.Set(ctx, token, constants.TOKEN_BLACKLIST_KEY, expiry)
}

func (b *Blacklist) IsTokenBlacklisted(ctx context.Context, token string) (bool, error) {
	return b.storeClient.SetIsMember(ctx, constants.TOKEN_BLACKLIST_SKEY, token)
}
