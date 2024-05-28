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

type IntermediateDatabaseSettings struct {
	Username     string `env:"USERNAME"`
	Password     string `env:"PASSWORD"`
	Port         uint   `env:"PORT"`
	Host         string `env:"HOST"`
	DatabaseName string `env:"NAME"`
	RequireSSL   bool   `env:"REQUIRE_SSL"`
}

func (i *IntermediateDatabaseSettings) Into() (*DatabaseSettings, error) {
	password, err := m.NewSecret(i.Password)
	if err != nil {
		return nil, err
	}

	return &DatabaseSettings{
		Username:     i.Username,
		Password:     password,
		Port:         i.Port,
		Host:         i.Host,
		DatabaseName: i.DatabaseName,
		RequireSSL:   i.RequireSSL,
	}, nil
}
