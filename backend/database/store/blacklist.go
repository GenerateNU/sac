package store

import "time"

type BlacklistInterface interface {
	BlacklistToken(token string, expiry time.Duration) error
	IsTokenBlacklisted(token string) (bool, error)
}

type Blacklist struct {
	StoreClient StoreClientInterface
}

func NewBlacklist(storeClient StoreClientInterface) *Blacklist {
	return &Blacklist{
		StoreClient: storeClient,
	}
}

func (b *Blacklist) BlacklistToken(token string, expiry time.Duration) error {
	err := b.StoreClient.SAdd("token_blacklist", token)
	if err != nil {
		return err
	}

	return b.StoreClient.Set(token, "blacklisted", expiry)
}

func (b *Blacklist) IsTokenBlacklisted(token string) (bool, error) {
	return b.StoreClient.SIsMember("token_blacklist", token)
}
