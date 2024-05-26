package config

import (
	"fmt"
	"time"

	m "github.com/garrettladley/mattress"
)

type Settings struct {
	Application ApplicationSettings
	Database    DatabaseSettings
	SuperUser   SuperUserSettings
	Auth        AuthSettings
	AWS         AWSSettings
	Resend      ResendSettings
	Calendar    CalendarSettings
}

type ApplicationSettings struct {
	Port    uint16
	Host    string
	BaseUrl string
}

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
		s.Host, s.Port, s.Username, s.Password.Expose(), sslMode)
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

	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", s.Username, s.Password.Expose(), s.Host, s.Port, s.DatabaseName, sslMode)
}

type SuperUserSettings struct {
	Password *m.Secret[string]
}

type AuthSettings struct {
	AccessKey  *m.Secret[string]
	RefreshKey *m.Secret[string]
}

type AWSSettings struct {
	BucketName *m.Secret[string]
	Id         *m.Secret[string]
	Secret     *m.Secret[string]
	Region     *m.Secret[string]
}

type ResendSettings struct {
	APIKey *m.Secret[string]
}

type CalendarSettings struct {
	MaxTerminationDate time.Time
}
