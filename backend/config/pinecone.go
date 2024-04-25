package config

import (
	"errors"
	"os"

	m "github.com/garrettladley/mattress"
)

type PineconeSettings struct {
	IndexHost *m.Secret[string]
	APIKey    *m.Secret[string]
}

func readPineconeSettings() (*PineconeSettings, error) {
	indexHost := os.Getenv("SAC_PINECONE_INDEX_HOST")
	if indexHost == "" {
		return nil, errors.New("SAC_PINECONE_INDEX_HOST is not set")
	}

	secretIndexHost, err := m.NewSecret(indexHost)
	if err != nil {
		return nil, errors.New("failed to create secret from index host")
	}

	apiKey := os.Getenv("SAC_PINECONE_API_KEY")
	if apiKey == "" {
		return nil, errors.New("SAC_PINECONE_API_KEY is not set")
	}

	secretAPIKey, err := m.NewSecret(apiKey)
	if err != nil {
		return nil, errors.New("failed to create secret from api key")
	}

	return &PineconeSettings{
		IndexHost: secretIndexHost,
		APIKey:    secretAPIKey,
	}, nil
}
