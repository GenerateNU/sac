package utilities

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/types"
)

var (
	indexHost = os.Getenv("SAC_PINECONE_INDEX_HOST")
	apiKey    = os.Getenv("SAC_PINECONE_API_KEY")
)

func UpsertToPinecone(item types.Embeddable) *errors.Error {
	embeddingResult, _err := item.Embed()
	if _err != nil {
		return &errors.FailedToUpsertPinecone
	}

	upsertBody, _ := json.Marshal(map[string]interface{}{
		"vectors": []types.Embedding{
			*embeddingResult,
		},
		"namespace": item.Namespace(),
	})
	requestBody := bytes.NewBuffer(upsertBody)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/vectors/upsert", indexHost), requestBody)
	if err != nil {
		return &errors.FailedToUpsertPinecone
	}

	req.Header.Set("Api-Key", apiKey)
	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &errors.FailedToUpsertPinecone
	}

	if resp.StatusCode != 200 {
		return &errors.FailedToUpsertPinecone
	}

	return nil
}

func DeleteFromPinecone(item types.Embeddable) *errors.Error {
	deleteBody, err := json.Marshal(map[string]interface{}{
		"deleteAll": false,
		"ids": []string{
			item.EmbeddingId(),
		},
		"namespace": item.Namespace(),
	})
	if err != nil {
		return &errors.FailedToDeletePinecone
	}
	requestBody := bytes.NewBuffer(deleteBody)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/vectors/delete", indexHost), requestBody)
	if err != nil {
		return &errors.FailedToDeletePinecone
	}

	req.Header.Set("Api-Key", apiKey)
	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &errors.FailedToDeletePinecone
	}

	if resp.StatusCode != 200 {
		return &errors.FailedToDeletePinecone
	}

	return nil
}

func SearchPinecone(item types.Embeddable, topKResults int) ([]string, *errors.Error) {
	embeddingResult, _err := item.Embed()
	if _err != nil {
		return []string{}, _err
	}

	searchBody, _ := json.Marshal(map[string]interface{}{
		"includeValues":   false,
		"includeMetadata": false,
		"topK":            topKResults,
		"vector":          embeddingResult.Values,
		"namespace":       item.Namespace(),
	})

	requestBody := bytes.NewBuffer(searchBody)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/query", indexHost), requestBody)
	if err != nil {
		return []string{}, &errors.FailedToSearchPinecone
	}

	req.Header.Set("Api-Key", apiKey)
	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return []string{}, &errors.FailedToSearchPinecone
	}

	if resp.StatusCode != 200 {
		return []string{}, &errors.FailedToSearchPinecone
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
		return []string{}, &errors.FailedToSearchPinecone
	}

	var resultsToReturn []string
	for i := 0; i < len(results.Matches); i += 1 {
		resultsToReturn = append(resultsToReturn, results.Matches[i].Id)
	}

	return resultsToReturn, nil
}
