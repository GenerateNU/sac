package redis

import (
	"time"

	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/gofiber/fiber/v2"
	goredis "github.com/redis/go-redis/v9"
)

type RedisClientInterface interface {
	BlacklistToken(c *fiber.Ctx, tokenString string, tokenType auth.JWTType) *errors.Error
	IsBlackListed(c *fiber.Ctx, tokenString string, tokenType auth.JWTType) (bool, *errors.Error)
}

type RedisClient struct {
	r *goredis.Client
	j auth.JWTClientInterface
}

func NewRedisClient(rs config.RedisSettings, jwtClient auth.JWTClientInterface) *RedisClient {
	redisClient := goredis.NewClient(
		&goredis.Options{
			Addr:     rs.Addr,
			Password: rs.Password,
			DB:       rs.DB,
		})

	return &RedisClient{
		r: redisClient,
		j: jwtClient,
	}
}

// checks if a token is blacklisted for a user
func (r *RedisClient) IsBlackListed(c *fiber.Ctx, tokenString string, tokenType auth.JWTType) (bool, *errors.Error) {

	tokenIsValid, err := r.j.IsTokenValid(tokenString, tokenType)

	if !tokenIsValid {
		return false, &errors.InvalidToken

	}

	if err != nil {
		return false, err
	}

	// Check if this token is in the blacklist
	exists, err1 := r.r.Exists(c.Context(), tokenString).Result()

	if err1 != nil {
		return false, &errors.Error{Message: err1.Error()}
	}

	return exists == 1, nil
}

// adds a token to the blacklist set for a user
func (r *RedisClient) BlacklistToken(c *fiber.Ctx, tokenString string, tokenType auth.JWTType) *errors.Error {

	tokenIsValid, err := r.j.IsTokenValid(tokenString, tokenType)

	if !tokenIsValid {
		return &errors.InvalidToken
	} else if err != nil {
		return err
	}

	claims, err := r.j.ExtractClaims(tokenString, tokenType)

	if err != nil {
		return err
	}

	expValue, ok := claims["exp"].(float64)

	if !ok {
		return &errors.FailedToGetExp
	}

	expTime := time.Unix(int64(expValue), 0)
	duration := time.Until(expTime)

	// set the token as a key in redis
	err1 := r.r.Set(c.Context(), tokenString, "blacklisted", duration).Err()

	if err1 != nil {
		return &errors.FailedToBlacklistToken
	}
	return nil
}
