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
	username string `env:"USERNAME"`
	password string `env:"PASSWORD"`
	host     string `env:"HOST"`
	port     uint   `env:"PORT"`
	db       int    `env:"DB"`
	// TLSConfig *intermediateTLSConfig `env:"TLS_CONFIG"`
}

func (i *intermediateRedisSettings) into() (*RedisSettings, error) {
	password, err := m.NewSecret(i.password)
	if err != nil {
		return nil, fmt.Errorf("failed to create secret from password: %s", err.Error())
	}

	return &RedisSettings{
		Username: i.username,
		Password: password,
		Host:     i.host,
		Port:     i.port,
		DB:       i.db,
		// TLSConfig: i.TLSConfig.into(),
	}, nil
}
