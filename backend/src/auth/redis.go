package auth

import (
	"github.com/gofiber/fiber/v2"

	"github.com/GenerateNU/sac/backend/src/errors"

	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/redis/go-redis/v9"
)

type RedisClientInterface interface {
	ValidateToken(key string) *errors.Error
	IsBlackListed(key string) (bool, *errors.Error)
}

type RedisClient struct {
	Client *redis.Client
	c   fiber.Ctx
}

func NewRedisClient(rs config.RedisSettings) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     rs.Addr,
		Password: rs.Password,
		DB:       rs.DB,
	})
}

func (r *RedisClient) ValidateToken(token string) *errors.Error {
	r.Client.LRange(r.c, "token",0,1)
	return nil
}

func (r *RedisClient) IsBlackListed(key string) (bool, *errors.Error) {

	return false, nil
}
