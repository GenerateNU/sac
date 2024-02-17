package config

import (
	"os"

	"github.com/spf13/viper"
)

type Settings struct {
	Application      ApplicationSettings
	Database         DatabaseSettings
	SuperUser        SuperUserSettings
	Auth             AuthSettings
	PineconeSettings PineconeSettings
	OpenAISettings   OpenAISettings
}

type intermediateSettings struct {
	Application ApplicationSettings           `yaml:"application"`
	Database    intermediateDatabaseSettings  `yaml:"database"`
	SuperUser   intermediateSuperUserSettings `yaml:"superuser"`
	Auth        intermediateAuthSettings      `yaml:"authsecret"`
}

func (int *intermediateSettings) into() (*Settings, error) {
	databaseSettings, err := int.Database.into()
	if err != nil {
		return nil, err
	}

	superUserSettings, err := int.SuperUser.into()
	if err != nil {
		return nil, err
	}

	authSettings, err := int.Auth.into()
	if err != nil {
		return nil, err
	}

	return &Settings{
		Application: int.Application,
		Database:    *databaseSettings,
		SuperUser:   *superUserSettings,
		Auth:        *authSettings,
	}, nil
}

type Environment string

const (
	EnvironmentLocal      Environment = "local"
	EnvironmentProduction Environment = "production"
)

func GetConfiguration(path string) (*Settings, error) {
	var environment Environment
	if env := os.Getenv("APP_ENVIRONMENT"); env != "" {
		environment = Environment(env)
	} else {
		environment = "local"
	}

	v := viper.New()
	v.SetConfigType("yaml")
	v.AddConfigPath(path)

	if environment == EnvironmentLocal {
		return readLocal(v, path)
	} else {
		return readProd(v)
	}
}
