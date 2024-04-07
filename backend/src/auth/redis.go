package auth

import (
	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type RedisClientInterface interface {
	BlacklistToken(c *fiber.Ctx, userId string, token string) *errors.Error
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

// checks if a token is blacklisted for a user
func (r *RedisClient) IsBlackListed(c *fiber.Ctx, userId string, token string) (bool, *errors.Error) {
	blacklistKey := UserBlacklistKey(userId)
	result, err := r.Client.SIsMember(c.Context(), blacklistKey, token).Result()
	if err != nil {
		return false, &errors.FailedToFindUserTokenList
	}

	return result, nil
}

// adds a token to the blacklist set for a user
func (r *RedisClient) BlacklistToken(c *fiber.Ctx, userId string, token string) *errors.Error {
	blacklistKey := UserBlacklistKey(userId)
	// maybe add expiration time to the key?
	numAdded, err := r.Client.SAdd(c.Context(), blacklistKey, token).Result()
	if err != nil || numAdded != 1 {
		return &errors.FailedToBlacklistToken
	}
	return nil
}

// TODO q: now that we're using a set to hold the keys, should we still separate the keys by user? or can we just have one key for all blacklisted tokens?
var UserBlacklistKey = func(userId string) string {
	return "Blacklist_" + userId
}
