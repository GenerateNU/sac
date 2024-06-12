package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/GenerateNU/sac/backend/config"
	"github.com/redis/go-redis/v9"
)

type RedisCacher struct {
	rdb *redis.Client
}

func NewRedisCacher(settings config.RedisSettings) *RedisCacher {
	return &RedisCacher{
		rdb: settings.Into(),
	}
}

func (c *RedisCacher) Get(ctx context.Context, key string, q *Query[any]) (*Query[any], error) {
	res, err := c.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	if err := q.Unmarshal([]byte(res)); err != nil {
		return nil, err
	}

	return q, nil
}

func (c *RedisCacher) Store(ctx context.Context, key string, val *Query[any], ttl time.Duration) error {
	res, err := val.Marshal()
	if err != nil {
		return err
	}

	c.rdb.Set(ctx, key, res, ttl)
	return nil
}

func (c *RedisCacher) Invalidate(ctx context.Context) error {
	var (
		cursor uint64
		keys   []string
	)
	for {
		var (
			k   []string
			err error
		)
		k, cursor, err = c.rdb.Scan(ctx, cursor, fmt.Sprintf("%s*", IdentifierPrefix), 0).Result()
		if err != nil {
			return err
		}
		keys = append(keys, k...)
		if cursor == 0 {
			break
		}
	}

	if len(keys) > 0 {
		if _, err := c.rdb.Del(ctx, keys...).Result(); err != nil {
			return err
		}
	}
	return nil
}
