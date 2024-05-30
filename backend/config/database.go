package config

import (
	"fmt"

	m "github.com/garrettladley/mattress"
)

type DatabaseSettings struct {
	Username     string
	Password     *m.Secret[string]
	Port         uint
	Host         string
	DatabaseName string
	RequireSSL   bool
}

func (s *DatabaseSettings) WithoutDb() string {
	var sslMode string
	if s.RequireSSL {
		sslMode = "require"
	} else {
		sslMode = "disable"
	}

	return fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=%s",
		s.Host, s.Port, s.Username, (*s.Password).Expose(), sslMode)
}

func (s *DatabaseSettings) WithDb() string {
	return fmt.Sprintf("%s dbname=%s", s.WithoutDb(), s.DatabaseName)
}

func (s *DatabaseSettings) PostgresConn() string {
	var sslMode string
	if s.RequireSSL {
		sslMode = "require"
	} else {
		sslMode = "disable"
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", s.Username, (*s.Password).Expose(), s.Host, s.Port, s.DatabaseName, sslMode)
}

type intermediateDatabaseSettings struct {
	username     string `env:"USERNAME"`
	password     string `env:"PASSWORD"`
	port         uint   `env:"PORT"`
	host         string `env:"HOST"`
	databaseName string `env:"NAME"`
	requireSSL   bool   `env:"REQUIRE_SSL"`
}

func (i *intermediateDatabaseSettings) into() (*DatabaseSettings, error) {
	password, err := m.NewSecret(i.password)
	if err != nil {
		return nil, fmt.Errorf("failed to create secret from password: %s", err.Error())
	}

	return &DatabaseSettings{
		Username:     i.username,
		Password:     password,
		Port:         i.port,
		Host:         i.host,
		DatabaseName: i.databaseName,
		RequireSSL:   i.requireSSL,
	}, nil
}
