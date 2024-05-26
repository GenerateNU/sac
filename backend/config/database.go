package config

import (
	"fmt"

	m "github.com/garrettladley/mattress"
)

// This is defined seperately from the other types because it is the only one that has
// member functions
type DatabaseSettings struct {
	Username     string            `env:"USERNAME" envDefault:"postgres"`
	Password     *m.Secret[string] `env:"PASSWORD" envDefault:"password"`
	Port         uint              `env:"PORT" envDefault:"5432"`
	Host         string            `env:"HOST" envDefault:"127.0.0.1"`
	DatabaseName string            `env:"NAME" envDefault:"sac"`
	RequireSSL   bool              `env:"REQUIRE_SSL" envDefault:"false"`
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
