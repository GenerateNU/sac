package config

import (
	"fmt"

	m "github.com/garrettladley/mattress"
)

type ResendSettings struct {
	APIKey *m.Secret[string]
}

type intermediateResendSettings struct {
	apiKey string `env:"API_KEY"`
}

func (i *intermediateResendSettings) into() (*ResendSettings, error) {
	apiKey, err := m.NewSecret(i.apiKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create secret from API key: %s", err.Error())
	}

	return &ResendSettings{
		APIKey: apiKey,
	}, nil
}
