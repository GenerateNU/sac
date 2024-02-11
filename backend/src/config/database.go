package config

import (
	"errors"
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

func (int *intermediateDatabaseSettings) into() (*DatabaseSettings, error) {
	password, err := m.NewSecret(int.Password)
	if err != nil {
		return nil, errors.New("failed to create secret from password")
	}

	return &DatabaseSettings{
		Username:     int.Username,
		Password:     password,
		Port:         int.Port,
		Host:         int.Host,
		DatabaseName: int.DatabaseName,
		RequireSSL:   int.RequireSSL,
	}, nil
}

type intermediateDatabaseSettings struct {
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Port         uint   `yaml:"port"`
	Host         string `yaml:"host"`
	DatabaseName string `yaml:"databasename"`
	RequireSSL   bool   `yaml:"requiressl"`
}

func (s *DatabaseSettings) WithoutDb() string {
	var sslMode string
	if s.RequireSSL {
		sslMode = "require"
	} else {
		sslMode = "disable"
	}

	return fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=%s",
		s.Host, s.Port, s.Username, s.Password.Expose(), sslMode)
}

func (s *DatabaseSettings) WithDb() string {
	return fmt.Sprintf("%s dbname=%s", s.WithoutDb(), s.DatabaseName)
}
