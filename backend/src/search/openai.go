package search

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/garrettladley/mattress"
	"net/http"
	"os"
)

type OpenAiClientInterface interface {
	CreateEmbedding(payload string) ([]float32, *errors.Error)
}

type OpenAiClient struct {
	apiKey *mattress.Secret[string]
}

func NewOpenAiClient() *OpenAiClient {
	apiKey, _ := mattress.NewSecret(os.Getenv("SAC_OPENAI_API_KEY"))

	return &OpenAiClient{apiKey: apiKey}
}

func (c *OpenAiClient) CreateEmbedding(payload string) ([]float32, *errors.Error) {
	apiKey := c.apiKey.Expose()

	embeddingBody, _ := json.Marshal(map[string]interface{}{
		"input": payload,
		"model": "text-embedding-ada-002",
	})
	requestBody := bytes.NewBuffer(embeddingBody)

	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.openai.com/v1/embeddings"), requestBody)
	if err != nil {
		return nil, &errors.FailedToCreateEmbedding
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("content-type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, &errors.FailedToCreateEmbedding
	}

	defer resp.Body.Close()

	if err != nil {
		return nil, &errors.FailedToCreateEmbedding
	}

	type ResponseBody struct {
		Data []struct {
			Embedding []float32 `json:"embedding"`
		} `json:"data"`
	}

	embeddingResultBody := ResponseBody{}
	err = json.NewDecoder(resp.Body).Decode(&embeddingResultBody)
	if err != nil {
		return nil, &errors.FailedToCreateEmbedding
	}

	if len(embeddingResultBody.Data) < 1 {
		return nil, &errors.FailedToCreateEmbedding
	}

	return embeddingResultBody.Data[0].Embedding, nil

}
