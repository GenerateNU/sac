package cache

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// SetUseCache enables or disables cache based on the value of `useCache`
func SetUseCache(db *gorm.DB, useCache bool) *gorm.DB {
	return db.Session(&gorm.Session{Context: setUseCache(useCache)})
}

// SetUseCacheTTL sets the cache TTL and ensures caching is enabled
func SetUseCacheTTL(db *gorm.DB, ttl time.Duration) *gorm.DB {
	return db.Session(&gorm.Session{Context: context.WithValue(setUseCache(true), cacheTTLKey, ttl)})
}

func setUseCache(useCache bool) context.Context {
	return context.WithValue(context.Background(), useCacheKey, useCache)
}
