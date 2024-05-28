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

	var intermediateSettings intermediateSettings
	if err := env.Parse(&intermediateSettings); err != nil {
		return nil, err
	}

	settings, err := intermediateSettings.into()
	return settings, err
}
