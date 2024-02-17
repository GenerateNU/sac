package config

import (
	"errors"
	"os"

	m "github.com/garrettladley/mattress"
)

type OpenAISettings struct {
	APIKey *m.Secret[string]
}

func readOpenAISettings() (*OpenAISettings, error) {
	apiKey := os.Getenv("SAC_OPENAI_API_KEY")
	if apiKey == "" {
		return nil, errors.New("SAC_OPENAI_API_KEY is not set")
	}

	secretAPIKey, err := m.NewSecret(apiKey)
	if err != nil {
		return nil, errors.New("failed to create secret from api key")
	}

	return &OpenAISettings{
		APIKey: secretAPIKey,
	}, nil
}
