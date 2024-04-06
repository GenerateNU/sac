package auth

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/redis/go-redis/v9"
	"time"

	"context"
	"git"

	"github.com/GenerateNU/sac/backend/src/config"
)

var (
	ctx = context.Background()
)

type RedisClientInterface interface {
	BlackListToken(userId string, token string, expiration time.Duration) *errors.Error
	IsBlackListed(userId string, token string) (bool, *errors.Error)
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

func (r *RedisClient) IsBlackListed(userId string, token string) (bool, *errors.Error) {
	blacklistKey := "Blacklist:" + userId
	tokenList, err := r.Client.LRange(*r.Client, ctx, blacklistKey, 0, -1).Result()
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

func (r *RedisClient) BlackListToken(userId string, token string, expiration time.Duration) *errors.Error {
	blacklistKey := "Blacklist:" + userId
	err := r.Client.RPush(*r.Client, ctx, blacklistKey, token, expiration)
	if err != nil {

	}
	return nil
}
