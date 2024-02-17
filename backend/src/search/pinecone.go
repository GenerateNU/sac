package search

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"

	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/utilities"
)

type PineconeClientInterface interface {
	Upsert(item Searchable) *errors.Error
	Delete(item Searchable) *errors.Error
	Search(item Searchable, topK int) ([]string, *errors.Error)
}

type PineconeClient struct {
	Settings     config.PineconeSettings
	openAIClient *OpenAIClient
}

func NewPineconeClient(openAIClient *OpenAIClient, settings config.PineconeSettings) *PineconeClient {
	return &PineconeClient{
		Settings:     settings,
		openAIClient: openAIClient,
	}
}

func (c *PineconeClient) pineconeRequest(req *http.Request) *http.Request {
	return utilities.ApplyModifiers(req,
		utilities.HeaderKV("Api-Key", c.Settings.APIKey.Expose()),
		utilities.AcceptJSON(),
		utilities.JSON(),
	)
}

type Vector struct {
	ID     string    `json:"id"`
	Values []float32 `json:"values"`
}

type PineconeUpsertRequestBody struct {
	Vectors   []Vector `json:"vectors"`
	Namespace string   `json:"namespace"`
}

func (c *PineconeClient) Upsert(item Searchable) *errors.Error {
	values, embeddingErr := c.openAIClient.CreateEmbedding(item.EmbeddingString())
	if embeddingErr != nil {
		return &errors.FailedToUpsertToPinecone
	}

	upsertBody, _ := json.Marshal(
		PineconeUpsertRequestBody{
			Vectors: []Vector{
				{
					ID:     item.SearchId(),
					Values: values,
				},
			},
			Namespace: item.Namespace(),
		})

	req, err := http.NewRequest(fiber.MethodPost,
		fmt.Sprintf("%s/vectors/upsert", c.Settings.IndexHost.Expose()),
		bytes.NewBuffer(upsertBody))
	if err != nil {
		return &errors.FailedToUpsertToPinecone
	}

	req = c.pineconeRequest(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &errors.FailedToUpsertToPinecone
	}

	if resp.StatusCode != fiber.StatusOK {
		return &errors.FailedToUpsertToPinecone
	}

	return nil
}

type PineconeDeleteRequestBody struct {
	IDs       []string `json:"ids"`
	Namespace string   `json:"namespace"`
	DeleteAll bool     `json:"deleteAll"`
}

func NewPineconeDeleteRequestBody(ids []string, namespace string, deleteAll bool) *PineconeDeleteRequestBody {
	return &PineconeDeleteRequestBody{
		IDs:       ids,
		Namespace: namespace,
		DeleteAll: deleteAll,
	}
}

func (c *PineconeClient) Delete(item Searchable) *errors.Error {
	deleteBody, err := json.Marshal(
		PineconeDeleteRequestBody{
			IDs:       []string{item.SearchId()},
			Namespace: item.Namespace(),
			DeleteAll: false,
		})
	if err != nil {
		return &errors.FailedToDeleteToPinecone
	}

	req, err := http.NewRequest(fiber.MethodPost,
		fmt.Sprintf("%s/vectors/delete", c.Settings.IndexHost.Expose()),
		bytes.NewBuffer(deleteBody))
	if err != nil {
		return &errors.FailedToDeleteToPinecone
	}

	req = c.pineconeRequest(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &errors.FailedToDeleteToPinecone
	}

	if resp.StatusCode != fiber.StatusOK {
		return &errors.FailedToDeleteToPinecone
	}

	return nil
}

type PineconeSearchRequestBody struct {
	IncludeValues   bool      `json:"includeValues"`
	IncludeMetadata bool      `json:"includeMetadata"`
	TopK            int       `json:"topK"`
	Vector          []float32 `json:"vector"`
	Namespace       string    `json:"namespace"`
}

type PineconeSearchResponseBody struct {
	Matches []struct {
		Id     string    `json:"id"`
		Score  float32   `json:"score"`
		Values []float32 `json:"values"`
	} `json:"matches"`
	Namespace string `json:"namespace"`
}

func (c *PineconeClient) Search(item Searchable, topK int) ([]string, *errors.Error) {
	values, embeddingErr := c.openAIClient.CreateEmbedding(item.EmbeddingString())
	if embeddingErr != nil {
		return []string{}, embeddingErr
	}

	searchBody, err := json.Marshal(
		PineconeSearchRequestBody{
			IncludeValues:   false,
			IncludeMetadata: false,
			TopK:            topK,
			Vector:          values,
			Namespace:       item.Namespace(),
		})
	if err != nil {
		return []string{}, &errors.FailedToSearchToPinecone
	}

	req, err := http.NewRequest(fiber.MethodPost,
		fmt.Sprintf("%s/query", c.Settings.IndexHost.Expose()),
		bytes.NewBuffer(searchBody))
	if err != nil {
		return []string{}, &errors.FailedToSearchToPinecone
	}

	req = c.pineconeRequest(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []string{}, &errors.FailedToSearchToPinecone
	}

	defer resp.Body.Close()

	if err != nil {
		return []string{}, &errors.FailedToSearchToPinecone
	}

	if resp.StatusCode != fiber.StatusOK {
		return []string{}, &errors.FailedToSearchToPinecone
	}

	var results PineconeSearchResponseBody
	err = json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		return []string{}, &errors.FailedToSearchToPinecone
	}

	resultsToReturn := make([]string, len(results.Matches))

	for i, match := range results.Matches {
		resultsToReturn[i] = match.Id
	}

	return resultsToReturn, nil
}
