package search

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"

	"github.com/GenerateNU/sac/backend/config"

	"github.com/GenerateNU/sac/backend/utilities"
)

type AIClientInterface interface {
	CreateEmbedding(items []Searchable) ([]Embedding, error)
	CreateModeration(items []Searchable) ([]ModerationResult, error)
}

type OpenAIClient struct {
	Settings config.OpenAISettings
}

func NewOpenAIClient(settings config.OpenAISettings) AIClientInterface {
	return &OpenAIClient{Settings: settings}
}

type CreateEmbeddingRequestBody struct {
	Input []string `json:"input"`
	Model string   `json:"model"`
}

type Embedding struct {
	Embedding []float32 `json:"embedding"`
}

type CreateEmbeddingResponseBody struct {
	Data []Embedding `json:"data"`
}

func (c *OpenAIClient) CreateEmbedding(items []Searchable) ([]Embedding, error) {
	embeddingStrings := make([]string, len(items))
	for i, item := range items {
		embeddingStrings[i] = item.EmbeddingString()
	}

	embeddingBody, err := json.Marshal(
		CreateEmbeddingRequestBody{
			Input: embeddingStrings,
			Model: "text-embedding-ada-002",
		})
	if err != nil {
		return nil, fmt.Errorf("failed to create embedding request body: %w", err)
	}

	resp, err := utilities.Request(http.MethodPost, "https://api.openai.com/v1/embeddings", embeddingBody, openAIModifiers(c.Settings)...)
	if err != nil {
		return nil, fmt.Errorf("failed to create embedding: %w", err)
	}

	var embeddingResultBody CreateEmbeddingResponseBody

	err = json.NewDecoder(resp.Body).Decode(&embeddingResultBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create embedding: %w", err)
	}

	if len(embeddingResultBody.Data) < 1 {
		return nil, fmt.Errorf("failed to create embedding: %w", err)
	}

	return embeddingResultBody.Data, nil
}

type CreateModerationRequestBody struct {
	Input []string `json:"input"`
	Model string   `json:"model"`
}

type CreateModerationResponseBody struct {
	Results []ModerationResult `json:"results"`
}

type ModerationResult struct {
	Flagged bool `json:"flagged"`
}

func (c *OpenAIClient) CreateModeration(items []Searchable) ([]ModerationResult, error) {
	searchStrings := make([]string, len(items))
	for i, item := range items {
		searchStrings[i] = item.EmbeddingString()
	}

	moderationBody, err := json.Marshal(
		CreateModerationRequestBody{
			Input: searchStrings,
			Model: "text-moderation-stable",
		})
	if err != nil {
		return nil, fmt.Errorf("failed to create moderation request body: %w", err)
	}

	resp, err := utilities.Request(http.MethodPost, "https://api.openai.com/v1/moderations", moderationBody, openAIModifiers(c.Settings)...)
	if err != nil {
		return nil, fmt.Errorf("failed to create moderation: %w", err)
	}

	var moderationResultBody CreateModerationResponseBody

	err = json.NewDecoder(resp.Body).Decode(&moderationResultBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create moderation: %w", err)
	}

	if len(moderationResultBody.Results) < 1 {
		return nil, errors.New("failed to create moderation")
	}

	return moderationResultBody.Results, nil
}

func openAIModifiers(settings config.OpenAISettings) []utilities.RequestModifier {
	return []utilities.RequestModifier{
		utilities.Authorization(settings.APIKey.Expose()),
		utilities.JSON(),
	}
}
