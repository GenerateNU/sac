package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

func GetConfiguration(path string) (*Settings, error) {
	if path != "" {
		if err := godotenv.Load(path); err != nil {
			return nil, fmt.Errorf("failed to load environment variables: %s", err.Error())
		}
	}

	intSettings, err := env.ParseAs[intermediateSettings]()
	if err != nil {
		return nil, fmt.Errorf("failed to parse environment variables: %s", err.Error())
	}

	return intSettings.into()
}
