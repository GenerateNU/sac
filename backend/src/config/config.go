package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Settings struct {
	Application ApplicationSettings `yaml:"application"`
	Database    DatabaseSettings    `yaml:"database"`
	SuperUser   SuperUserSettings   `yaml:"superuser"`
	Auth        AuthSettings
	AWS         AWSSettings
	PineconeSettings PineconeSettings
	OpenAISettings   OpenAISettings
}

type intermediateSettings struct {
	Application ApplicationSettings           `yaml:"application"`
	Database    intermediateDatabaseSettings  `yaml:"database"`
	SuperUser   intermediateSuperUserSettings `yaml:"superuser"`
	Auth        intermediateAuthSettings      `yaml:"authsecret"`
	AWS         AWSSettings
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

func configAWS() AWSSettings {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	return AWSSettings{
		BUCKET_NAME: os.Getenv("BUCKET_NAME"),
		ID:          os.Getenv("AWS_ID"),
		SECRET:      os.Getenv("AWS_SECRET")}
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
