package mocks

import (
	"github.com/GenerateNU/sac/backend/integrations/search"
	"gorm.io/gorm"
)

type PineconeMockClient struct{}

// Connects to an existing Pinecone index, using the host and keys provided in settings.
func NewPineconeMockClient() search.SearchClientInterface {
	return &PineconeMockClient{}
}

func (c *PineconeMockClient) Upsert(items []search.Searchable) error {
	return nil
}

func (c *PineconeMockClient) Delete(items []search.Searchable) error {
	return nil
}

func (c *PineconeMockClient) Search(item search.Searchable) ([]string, error) {
	return []string{}, nil
}

func (c *PineconeMockClient) Seed(db *gorm.DB) error {
	return nil
}
