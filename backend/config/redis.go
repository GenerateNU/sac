package config

import (
	"errors"

	m "github.com/garrettladley/mattress"
)

type RedisSettings struct {
	Username string
	Password *m.Secret[string]
	Host     string
	Port     uint
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
		Password: password,
		Host:     int.Host,
		Port:     int.Port,
		DB:       int.DB,
		// TLSConfig: int.TLSConfig.into(),
	}, nil
}

type intermediateRedisSettings struct {
	Username string `env:"USERNAME"`
	Password string `env:"PASSWORD"`
	Host     string `env:"HOST"`
	Port     uint   `env:"PORT"`
	DB       int    `env:"DB"`
	// TLSConfig *intermediateTLSConfig `env:"TLS_CONFIG"`
}
