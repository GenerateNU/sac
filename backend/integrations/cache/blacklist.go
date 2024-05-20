package cache

import "time"

type BlacklistInterface interface {
	BlacklistToken(token string, ttl time.Duration) error
	IsTokenBlacklisted(token string) (bool, error)
}

type Blacklist struct {
	CacheClient CacheClientInterface
}

func NewBlacklist(cacheClient CacheClientInterface) *Blacklist {
	return &Blacklist{
		CacheClient: cacheClient,
	}
}

func (b *Blacklist) BlacklistToken(token string, ttl time.Duration) error {
	err := b.CacheClient.SAdd("token_blacklist", token)
	if err != nil {
		return err
	}

	return b.CacheClient.Set(token, "blacklisted", ttl)
}

func (b *Blacklist) IsTokenBlacklisted(token string) (bool, error) {
	return b.CacheClient.Exists(token)
}

func (b *Blacklist) IsRefreshTokenBlacklisted(token string) (bool, error) {
	return b.CacheClient.SIsMember("token_blacklist", token)
}