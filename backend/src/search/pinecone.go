package search

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/garrettladley/mattress"
	"net/http"
	"os"

	"github.com/GenerateNU/sac/backend/src/errors"
)

type PineconeClientInterface interface {
	Upsert(item Searchable) *errors.Error
	Delete(item Searchable) *errors.Error
	Search(item Searchable, topK int) ([]string, *errors.Error)
}

type PineconeClient struct {
	indexHost    *mattress.Secret[string]
	apiKey       *mattress.Secret[string]
	openAiClient *OpenAiClient
}

func NewPineconeClient[T OpenAiClient](openAiClient *OpenAiClient) *PineconeClient {
	indexHost, _ := mattress.NewSecret(os.Getenv("SAC_PINECONE_INDEX_HOST"))
	apiKey, _ := mattress.NewSecret(os.Getenv("SAC_PINECONE_API_KEY"))

	return &PineconeClient{indexHost: indexHost, apiKey: apiKey, openAiClient: openAiClient}
}

func (c *PineconeClient) Upsert(item Searchable) *errors.Error {
	values, _err := c.openAiClient.CreateEmbedding(item.EmbeddingString())
	if _err != nil {
		return &errors.FailedToUpsertToPinecone
	}

	upsertBody, _ := json.Marshal(map[string]interface{}{
		"vectors": []interface{}{
			map[string]interface{}{
				"id":     item.SearchId(),
				"values": values,
			},
		},
		"namespace": item.Namespace(),
	})
	requestBody := bytes.NewBuffer(upsertBody)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/vectors/upsert", c.indexHost.Expose()), requestBody)
	if err != nil {
		return &errors.FailedToUpsertToPinecone
	}

	req.Header.Set("Api-Key", c.apiKey.Expose())
	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &errors.FailedToUpsertToPinecone
	}

	if resp.StatusCode != 200 {
		return &errors.FailedToUpsertToPinecone
	}

	return nil
}

func (c *PineconeClient) Delete(item Searchable) *errors.Error {
	deleteBody, err := json.Marshal(map[string]interface{}{
		"deleteAll": false,
		"ids": []string{
			item.SearchId(),
		},
		"namespace": item.Namespace(),
	})
	if err != nil {
		return &errors.FailedToDeleteToPinecone
	}
	requestBody := bytes.NewBuffer(deleteBody)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/vectors/delete", c.indexHost.Expose()), requestBody)
	if err != nil {
		return &errors.FailedToDeleteToPinecone
	}

	req.Header.Set("Api-Key", c.apiKey.Expose())
	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &errors.FailedToDeleteToPinecone
	}

	if resp.StatusCode != 200 {
		return &errors.FailedToDeleteToPinecone
	}

	return nil
}

func (c *PineconeClient) Search(item Searchable, topK int) ([]string, *errors.Error) {
	values, _err := c.openAiClient.CreateEmbedding(item.EmbeddingString())
	if _err != nil {
		return []string{}, _err
	}

	searchBody, _ := json.Marshal(map[string]interface{}{
		"includeValues":   false,
		"includeMetadata": false,
		"topK":            topK,
		"vector":          values,
		"namespace":       item.Namespace(),
	})
	requestBody := bytes.NewBuffer(searchBody)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/query", c.indexHost.Expose()), requestBody)
	if err != nil {
		return []string{}, &errors.FailedToSearchToPinecone
	}

	req.Header.Set("Api-Key", c.apiKey.Expose())
	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []string{}, &errors.FailedToSearchToPinecone
	}

	defer resp.Body.Close()

	if err != nil {
		return []string{}, &errors.FailedToSearchToPinecone
	}

	if resp.StatusCode != 200 {
		return []string{}, &errors.FailedToSearchToPinecone
	}

	type SearchPineconeResults struct {
		Matches []struct {
			Id     string    `json:"id"`
			Score  float32   `json:"score"`
			Values []float32 `json:"values"`
		} `json:"matches"`
		Namespace string `json:"namespace"`
	}

	results := SearchPineconeResults{}
	err = json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		return []string{}, &errors.FailedToSearchToPinecone
	}

	var resultsToReturn []string
	for i := 0; i < len(results.Matches); i += 1 {
		resultsToReturn = append(resultsToReturn, results.Matches[i].Id)
	}

	return resultsToReturn, nil
}
