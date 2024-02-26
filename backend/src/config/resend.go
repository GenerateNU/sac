package config

import (
	"errors"
	"os"

	m "github.com/garrettladley/mattress"
)

type ResendSettings struct {
	APIKey *m.Secret[string]
}

func readResendSettings() (*ResendSettings, error) {
	apiKey := os.Getenv("SAC_RESEND_API_KEY")
	if apiKey == "" {
		return nil, errors.New("SAC_RESEND_API_KEY is not set")
	}

	secretAPIKey, err := m.NewSecret(apiKey)
	if err != nil {
		return nil, errors.New("failed to create secret from api key")
	}

	return &ResendSettings{
		APIKey: secretAPIKey,
	}, nil
}