package search

import (
	"bytes"
	"io"
	"net/http"

	"github.com/goccy/go-json"

	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/gofiber/fiber/v2"
)

type OpenAIClientInterface interface {
	CreateEmbedding(payload string) ([]float32, *errors.Error)
}

type OpenAIClient struct {
	Settings config.OpenAISettings
}

func NewOpenAIClient(settings config.OpenAISettings) *OpenAIClient {
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

func (c *OpenAIClient) CreateEmbedding(items []Searchable) ([]Embedding, *errors.Error) {
	embeddingStrings := []string{}
	for _, item := range items {
		embeddingStrings = append(embeddingStrings, item.EmbeddingString())
	}

	embeddingBody, err := json.Marshal(
		CreateEmbeddingRequestBody{
			Input: embeddingStrings,
			Model: "text-embedding-ada-002",
		})
	if err != nil {
		return nil, &errors.FailedToCreateEmbedding
	}

	req, err := http.NewRequest(fiber.MethodPost,
		"https://api.openai.com/v1/embeddings",
		bytes.NewBuffer(embeddingBody))
	if err != nil {
		return nil, &errors.FailedToCreateEmbedding
	}

	req = utilities.ApplyModifiers(req,
		utilities.Authorization(c.Settings.APIKey.Expose()),
		utilities.JSON(),
	)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, &errors.FailedToCreateEmbedding
	}

	defer resp.Body.Close()

	var embeddingResultBody CreateEmbeddingResponseBody

	err = json.NewDecoder(resp.Body).Decode(&embeddingResultBody)
	if err != nil {
		return nil, &errors.FailedToCreateEmbedding
	}

	if len(embeddingResultBody.Data) < 1 {
		return nil, &errors.FailedToCreateEmbedding
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

func (c *OpenAIClient) CreateModeration(items []Searchable) ([]ModerationResult, *errors.Error) {
	searchStrings := []string{}
	for _, item := range items {
		searchStrings = append(searchStrings, item.EmbeddingString())
	}

	moderationBody, err := json.Marshal(
		CreateModerationRequestBody{
			Input: searchStrings,
			Model: "text-moderation-stable",
		})
	if err != nil {
		return nil, &errors.FailedToCreateModeration
	}

	req, err := http.NewRequest(fiber.MethodPost,
		"https://api.openai.com/v1/moderations",
		bytes.NewBuffer(moderationBody))
	if err != nil {
		return nil, &errors.FailedToCreateModeration
	}

	req = utilities.ApplyModifiers(req,
		utilities.Authorization(c.Settings.APIKey.Expose()),
		utilities.JSON(),
	)

	resp, err := http.DefaultClient.Do(req)

	respBody, err := io.ReadAll(resp.Body)
	respBodyString := string(respBody)
	print(respBodyString)

	if err != nil {
		return nil, &errors.FailedToCreateModeration
	}

	defer resp.Body.Close()

	var moderationResultBody CreateModerationResponseBody

	err = json.NewDecoder(resp.Body).Decode(&moderationResultBody)
	if err != nil {
		return nil, &errors.FailedToCreateModeration
	}

	if len(moderationResultBody.Results) < 1 {
		return nil, &errors.FailedToCreateModeration
	}

	return moderationResultBody.Results, nil
}
