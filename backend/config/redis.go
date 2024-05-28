package config

import (
	"errors"

	m "github.com/garrettladley/mattress"
)

type RedisSettings struct {
	Username string            `yaml:"username"`
	Host     string            `yaml:"host"`
	Port     uint              `yaml:"port"`
	Password *m.Secret[string] `yaml:"password"`
	DB       int               `yaml:"db"`
	// TLSConfig *TLSConfig       `yaml:"tlsconfig"`
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
	Username string `yaml:"username"`
	Host     string `yaml:"host"`
	Port     uint   `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
	// TLSConfig *intermediateTLSConfig `yaml:"tlsconfig"`
}
