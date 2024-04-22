package search

import "github.com/GenerateNU/sac/backend/errors"

type OpenAIMockClient struct{}

func NewOpenAIMockClient() AIClientInterface {
	return &OpenAIMockClient{}
}

func (c *OpenAIMockClient) CreateEmbedding(items []Searchable) ([]Embedding, *errors.Error) {
	return []Embedding{}, nil
}

func (c *OpenAIMockClient) CreateModeration(items []Searchable) ([]ModerationResult, *errors.Error) {
	return []ModerationResult{}, nil
}
