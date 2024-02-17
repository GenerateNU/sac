package search

import (
	"bytes"
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
	Input string `json:"input"`
	Model string `json:"model"`
}

type Embedding struct {
	Embedding []float32 `json:"embedding"`
}

type CreateEmbeddingResponseBody struct {
	Data []Embedding `json:"data"`
}

func (c *OpenAIClient) CreateEmbedding(payload string) ([]float32, *errors.Error) {
	embeddingBody, err := json.Marshal(
		CreateEmbeddingRequestBody{
			Input: payload,
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

	if err != nil {
		return nil, &errors.FailedToCreateEmbedding
	}

	var embeddingResultBody CreateEmbeddingResponseBody

	err = json.NewDecoder(resp.Body).Decode(&embeddingResultBody)
	if err != nil {
		return nil, &errors.FailedToCreateEmbedding
	}

	if len(embeddingResultBody.Data) < 1 {
		return nil, &errors.FailedToCreateEmbedding
	}

	EMBEDDING_INDEX := 0

	return embeddingResultBody.Data[EMBEDDING_INDEX].Embedding, nil
}
