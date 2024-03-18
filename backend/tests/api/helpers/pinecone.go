package helpers

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/search"
)

type PineconeMockClient struct{}

// Connects to an existing Pinecone index, using the host and keys provided in settings.
func NewPineconeMockClient() *PineconeMockClient {
	return &PineconeMockClient{}
}

func (c *PineconeMockClient) Upsert(items []search.Searchable) *errors.Error {
	return nil
}

func (c *PineconeMockClient) Delete(items []search.Searchable) *errors.Error {
	return nil
}

func (c *PineconeMockClient) Search(item search.Searchable, topK int) ([]string, *errors.Error) {
	return []string{}, nil
}
