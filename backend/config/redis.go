package config

import (
	"fmt"

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
		Username: i.Username,
		Password: password,
		Host:     i.Host,
		Port:     i.Port,
		DB:       i.DB,
		// TLSConfig: i.TLSConfig.into(),
	}, nil
}
