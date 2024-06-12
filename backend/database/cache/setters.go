package cache

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// SetCache enables or disables cache based on the value of `useCache`
func SetCache(db *gorm.DB, useCache bool) *gorm.DB {
	return db.Session(&gorm.Session{Context: setUseCache(useCache)})
}

// SetCacheTTL sets the cache TTL and ensures caching is enabled
func SetCacheTTL(db *gorm.DB, ttl time.Duration) *gorm.DB {
	return db.Session(&gorm.Session{Context: context.WithValue(setUseCache(true), cacheTTLKey, ttl)})
}

func setUseCache(useCache bool) context.Context {
	return context.WithValue(context.Background(), useCacheKey, useCache)
}
