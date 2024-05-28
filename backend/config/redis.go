package config

import (
	"errors"

	m "github.com/garrettladley/mattress"
)

type RedisSettings struct {
	Username string
	Host     string
	Port     uint
	Password *m.Secret[string]
	DB       int
	// TLSConfig *TLSConfig
}

func (int *intermediateRedisSettings) into() (*RedisSettings, error) {
	password, err := m.NewSecret(int.Password)
	if err != nil {
		return nil, errors.New("failed to create secret from password")
	}

	return &RedisSettings{
		Username: int.Username,
		Host:     int.Host,
		Port:     int.Port,
		Password: password,
		DB:       int.DB,
		// TLSConfig: int.TLSConfig.into(),
	}, nil
}

type intermediateRedisSettings struct {
	Username string `env:"USERNAME"`
	Host     string `env:"HOST"`
	Port     uint   `env:"PORT"`
	Password string `env:"PASSWORD"`
	DB       int    `env:"DB"`
	// TLSConfig *intermediateTLSConfig `env:"TLS_CONFIG"`
}
