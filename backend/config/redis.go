package config

import (
	"fmt"

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

func (r RedisSettings) Username() string {
	return r.username
}

func (r RedisSettings) Password() *m.Secret[string] {
	return r.password
}

func (r RedisSettings) Host() string {
	return r.host
}

func (r RedisSettings) Port() uint {
	return r.port
}

func (r RedisSettings) DB() int {
	return r.db
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
