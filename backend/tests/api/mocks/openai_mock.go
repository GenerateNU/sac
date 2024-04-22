package mocks

import (
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/GenerateNU/sac/backend/search"
)

type OpenAIMockClient struct{}

func NewOpenAIMockClient() search.AIClientInterface {
	return &OpenAIMockClient{}
}

func (c *OpenAIMockClient) CreateEmbedding(items []search.Searchable) ([]search.Embedding, *errors.Error) {
	return []search.Embedding{}, nil
}

func (c *OpenAIMockClient) CreateModeration(items []search.Searchable) ([]search.ModerationResult, *errors.Error) {
	return []search.ModerationResult{}, nil
}
