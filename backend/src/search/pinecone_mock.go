package search

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"gorm.io/gorm"
)

type PineconeMockClient struct{}

// Connects to an existing Pinecone index, using the host and keys provided in settings.
func NewPineconeMockClient() SearchClientInterface {
	return &PineconeMockClient{}
}

func (c *PineconeMockClient) Upsert(items []Searchable) *errors.Error {
	return nil
}

func (c *PineconeMockClient) Delete(items []Searchable) *errors.Error {
	return nil
}

func (c *PineconeMockClient) Search(item Searchable) ([]string, *errors.Error) {
	return []string{}, nil
}

func (c *PineconeMockClient) Seed(db *gorm.DB) *errors.Error {
	return nil
}
