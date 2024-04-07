package auth

import (
	"time"
	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type RedisClientInterface interface {
	BlacklistToken(c *fiber.Ctx, userId string, token string, expiration time.Duration) *errors.Error
	IsBlackListed(c *fiber.Ctx, userId string, token string) (bool, *errors.Error)
}

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(rs config.RedisSettings) *RedisClient {
	client := redis.NewClient(
		&redis.Options{
			Addr:     rs.Addr,
			Password: rs.Password,
			DB:       rs.DB,
		})

	return &RedisClient{
		Client: client,
	}
}

func (r *RedisClient) IsBlackListed(c *fiber.Ctx, userId string, token string) (bool, *errors.Error) {
	blacklistKey := "Blacklist_" + userId
	tokenList, err := r.Client.LRange(c.Context(), blacklistKey, 0, -1).Result()
	if err != nil {
		return false, &errors.FailedToFindUserTokenList
	}
	for _, tokenVal := range tokenList {
		if tokenVal == token {
			return true, nil
		}
	}
	return false, nil
}

func (r *RedisClient) BlacklistToken(c *fiber.Ctx, userId string, token string, expiration time.Duration) *errors.Error {
	blacklistKey := "Blacklist:" + userId
	err := r.Client.RPush(c.Context(), blacklistKey, token, expiration)
	if err != nil {
		return &errors.FailedToBlacklistToken
	}
	return nil
}
