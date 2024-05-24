package search

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"

	"github.com/GenerateNU/sac/backend/config"

	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
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

	req, err := http.NewRequest(fiber.MethodPost,
		"https://api.openai.com/v1/embeddings",
		bytes.NewBuffer(embeddingBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create embedding: %w", err)
	}

	req = utilities.ApplyModifiers(req,
		utilities.Authorization(c.Settings.APIKey.Expose()),
		utilities.JSON(),
	)

	resp, err := utilities.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create embedding: %w", err)
	}
	defer resp.Body.Close()

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

	req, err := http.NewRequest(fiber.MethodPost,
		"https://api.openai.com/v1/moderations",
		bytes.NewBuffer(moderationBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create moderation: %w", err)
	}

	req = utilities.ApplyModifiers(req,
		utilities.Authorization(c.Settings.APIKey.Expose()),
		utilities.JSON(),
	)

	resp, err := utilities.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create moderation: %w", err)
	}
	defer resp.Body.Close()

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
