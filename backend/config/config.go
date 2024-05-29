package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

func GetConfiguration(path string) (*Settings, error) {
	err := godotenv.Load(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load environment variables: %s", err.Error())
	}

	var intermediateSettings intermediateSettings
	if err := env.Parse(&intermediateSettings); err != nil {
		return nil, fmt.Errorf("failed to parse environment variables: %s", err.Error())
	}

	settings, err := intermediateSettings.into()
	return settings, err
}
