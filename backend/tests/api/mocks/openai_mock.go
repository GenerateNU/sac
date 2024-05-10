package mocks

import (
	"github.com/GenerateNU/sac/backend/integrations/search"
)

type OpenAIMockClient struct{}

func NewOpenAIMockClient() search.AIClientInterface {
	return &OpenAIMockClient{}
}

func (c *OpenAIMockClient) CreateEmbedding(items []search.Searchable) ([]search.Embedding, error) {
	return []search.Embedding{}, nil
}

func (c *OpenAIMockClient) CreateModeration(items []search.Searchable) ([]search.ModerationResult, error) {
	return []search.ModerationResult{}, nil
}
