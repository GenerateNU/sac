package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Settings struct {
	Application     ApplicationSettings
	Database        DatabaseSettings
	Redis       []RedisSettings
	SuperUser       SuperUserSettings
	Auth            AuthSettings
	AWS             AWSSettings
	Pinecone        PineconeSettings
	OpenAI          OpenAISettings
	Resend          ResendSettings
	Calendar        CalendarSettings
	GoogleSettings  OAuthSettings
	OutlookSettings OAuthSettings
}

type intermediateSettings struct {
	Application ApplicationSettings           `yaml:"application"`
	Database    intermediateDatabaseSettings  `yaml:"database"`
	Redis       []intermediateRedisSettings   `yaml:"redis"`
	SuperUser   intermediateSuperUserSettings `yaml:"superuser"`
	Auth        intermediateAuthSettings      `yaml:"authsecret"`
	Calendar    intermediateCalendarSettings  `yaml:"calendar"`
}

func (int *intermediateSettings) into() (*Settings, error) {
	databaseSettings, err := int.Database.into()
	if err != nil {
		return nil, err
	}

	redisSettings := make([]RedisSettings, len(int.Redis))
	for i, r := range int.Redis {
		redisInstance, err := r.into()
		if err != nil {
			return nil, err
		}

		redisSettings[i] = *redisInstance
	}

	superUserSettings, err := int.SuperUser.into()
	if err != nil {
		return nil, err
	}

	authSettings, err := int.Auth.into()
	if err != nil {
		return nil, err
	}

	calendarSettings, err := int.Calendar.into()
	if err != nil {
		return nil, err
	}

	return &Settings{
		Application: int.Application,
		Database:    *databaseSettings,
		Redis:       redisSettings,
		SuperUser:   *superUserSettings,
		Auth:        *authSettings,
		Calendar:    *calendarSettings,
	}, nil
}

type Environment string

const (
	EnvironmentLocal      Environment = "local"
	EnvironmentProduction Environment = "production"
)

func GetConfiguration(path string, useDevDotEnv bool) (*Settings, error) {
	var environment Environment
	if env := os.Getenv("APP_ENVIRONMENT"); env != "" {
		environment = Environment(env)
	} else {
		environment = "local"
	}

	v := viper.New()
	v.SetConfigType("yaml")
	v.AddConfigPath(path)

	switch environment {
	case EnvironmentLocal:
		return readLocal(v, path, useDevDotEnv)
	case EnvironmentProduction:
		return readProd(v)
	default:
		return nil, fmt.Errorf("unknown environment: %s", environment)
	}
}
