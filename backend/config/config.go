package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

func GetConfiguration(path string) (*Settings, error) {
	err := godotenv.Load(path)
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

	var intermediateSettings IntermediateSettings
	if err := env.Parse(&intermediateSettings); err != nil {
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
