package cache

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log/slog"
	"time"

	"github.com/GenerateNU/sac/backend/config"
	go_json "github.com/goccy/go-json"

	"gorm.io/gorm/callbacks"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type key byte

const (
	useCacheKey key = 0
	cacheTTLKey key = 1
)

// GormCache is a cache plugin for gorm
type GormCache struct {
	name   string
	client CacheClient
	config CacheConfig
}

// CacheClient is an interface for cache operations
type CacheClient interface {
	Get(ctx context.Context, key string) (interface{}, error)
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Delete(ctx context.Context, key string) error // Add Delete method
}

// CacheConfig is a struct for cache options
type CacheConfig struct {
	TTL    time.Duration // cache expiration time
	Prefix string        // cache key prefix
}

// NewGormCache returns a new GormCache instance
func NewGormCache(name string, client CacheClient, config CacheConfig) *GormCache {
	return &GormCache{
		name:   name,
		client: client,
		config: config,
	}
}

// Name returns the plugin name
func (g *GormCache) Name() string {
	return g.name
}

// Initialize initializes the plugin
func (g *GormCache) Initialize(db *gorm.DB) error {
	if err := db.Callback().Query().Replace("gorm:query", g.queryCallback); err != nil {
		return err
	}

	if err := db.Callback().Create().After("gorm:create").Register("cache:invalidate_create", g.invalidateCallback); err != nil {
		return err
	}

	if err := db.Callback().Update().After("gorm:update").Register("cache:invalidate_update", g.invalidateCallback); err != nil {
		return err
	}

	if err := db.Callback().Delete().After("gorm:delete").Register("cache:invalidate_delete", g.invalidateCallback); err != nil {
		return err
	}

	return nil
}

// queryCallback is a callback function for query operations
func (g *GormCache) queryCallback(db *gorm.DB) {
	if db.Error != nil {
		return
	}

	enableCache := g.enableCache(db)

	// build query sql
	callbacks.BuildQuerySQL(db)
	if db.DryRun || db.Error != nil {
		return
	}

	var (
		key string
		err error
		hit bool
	)
	if enableCache {
		key = g.cacheKey(db)

		// get value from cache
		hit, err = g.loadCache(db, key)
		if err != nil {
			slog.Error("load cache failed", "error", err, "hit", hit)
			return
		}

		// hit cache
		if hit {
			return
		}
	}

	if !hit {
		g.queryDB(db)

		if enableCache {
			if err = g.SetUseCache(db, key); err != nil {
				slog.Error("set cache failed", "error", err)
			}
		}
	}
}

func (g *GormCache) enableCache(db *gorm.DB) bool {
	ctx := db.Statement.Context

	// check if use cache
	useCache, ok := ctx.Value(useCacheKey).(bool)
	if !ok || !useCache {
		return false // do not use cache, skip this callback
	}
	return true
}

func (g *GormCache) cacheKey(db *gorm.DB) string {
	sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
	hash := sha256.Sum256([]byte(sql))
	key := g.config.Prefix + hex.EncodeToString(hash[:])
	return key
}

func (g *GormCache) loadCache(db *gorm.DB, key string) (bool, error) {
	value, err := g.client.Get(db.Statement.Context, key)
	if err != nil && !errors.Is(err, redis.Nil) {
		return false, err
	}

	if value == nil {
		return false, nil
	}

	// cache hit, scan value to destination
	if err = go_json.Unmarshal(value.([]byte), &db.Statement.Dest); err != nil {
		return false, err
	}
	db.RowsAffected = int64(db.Statement.ReflectValue.Len())
	return true, nil
}

func (g *GormCache) SetUseCache(db *gorm.DB, key string) error {
	ctx := db.Statement.Context

	// get cache ttl from context or config
	ttl, ok := ctx.Value(cacheTTLKey).(time.Duration)
	if !ok {
		ttl = g.config.TTL // use default ttl
	}

	// set value to cache with ttl
	return g.client.Set(ctx, key, db.Statement.Dest, ttl)
}

func (g *GormCache) queryDB(db *gorm.DB) {
	rows, err := db.Statement.ConnPool.QueryContext(db.Statement.Context, db.Statement.SQL.String(), db.Statement.Vars...)
	if err != nil {
		if err := db.AddError(err); err != nil {
			slog.Error("error encountered while adding error", "error", err)
		}
		return
	}
	defer func() {
		if err := db.AddError(rows.Close()); err != nil {
			slog.Error("error encountered while closing rows", "error", err)
		}
	}()
	gorm.Scan(rows, db, 0)
}

func (g *GormCache) invalidateCallback(db *gorm.DB) {
	if db.Error != nil {
		return
	}

	if err := g.invalidateCache(db); err != nil {
		slog.Error("invalidate cache failed", "error", err)
	}
}

func (g *GormCache) invalidateCache(db *gorm.DB) error {
	return g.client.Delete(db.Statement.Context, g.cacheKey(db))
}

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(settings config.RedisSettings) *RedisClient {
	return &RedisClient{
		client: settings.Into(),
	}
}

// Get gets value from redis by key using json encoding/decoding
func (r *RedisClient) Get(ctx context.Context, key string) (interface{}, error) {
	data, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Set sets value to redis by key with ttl using json encoding/decoding
func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	data, err := go_json.Marshal(value) // encode value to json bytes using json encoding/decoding
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, data, ttl).Err()
}

func (r *RedisClient) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}
