package config

import (
	"fmt"
	"runtime"

	"github.com/GenerateNU/sac/backend/constants"
	"github.com/redis/go-redis/v9"

	m "github.com/garrettladley/mattress"
)

type RedisSettings struct {
	username string
	password *m.Secret[string]
	host     string
	port     uint
	db       int
	// TLSConfig *TLSConfig
}

func (r *RedisSettings) Into() *redis.Client {
	return redis.NewClient(&redis.Options{
		Username:              r.username,
		Password:              r.password.Expose(),
		Addr:                  fmt.Sprintf("%s:%d", r.host, r.port),
		DB:                    r.db,
		PoolSize:              10 * runtime.GOMAXPROCS(0),
		MaxActiveConns:        constants.REDIS_MAX_OPEN_CONNECTIONS,
		MaxIdleConns:          constants.REDIS_MAX_IDLE_CONNECTIONS,
		ContextTimeoutEnabled: true,
	})
}

type intermediateRedisSettings struct {
	Username string `env:"USERNAME"`
	Password string `env:"PASSWORD"`
	Host     string `env:"HOST"`
	Port     uint   `env:"PORT"`
	DB       int    `env:"DB"`
	// TLSConfig *intermediateTLSConfig `env:"TLS_CONFIG"`
}

func (i *intermediateRedisSettings) into() (*RedisSettings, error) {
	password, err := m.NewSecret(i.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to create secret from password: %s", err.Error())
	}

	return &RedisSettings{
		username: i.Username,
		password: password,
		host:     i.Host,
		port:     i.Port,
		db:       i.DB,
		// TLSConfig: i.TLSConfig.into(),
	}, nil
}
