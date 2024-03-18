package config

import (
	"errors"
	"os"

	m "github.com/garrettladley/mattress"
)

type ClerkSettings struct {
	APIKey *m.Secret[string]
}

func readClerkSettings() (*ClerkSettings, error) {
	apiKey := os.Getenv("SAC_CLERK_SECRET_KEY")
	if apiKey == "" {
		return nil, errors.New("SAC_CLERK_SECRET_KEY is not set")
	}

	secretAPIKey, err := m.NewSecret(apiKey)
	if err != nil {
		return nil, errors.New("failed to create secret from api key")
	}

	return &ClerkSettings{
		APIKey: secretAPIKey,
	}, nil
}
