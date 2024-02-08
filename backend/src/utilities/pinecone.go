package utilities

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/types"
)

// TODO: move to config.yml files, then pass in to UpsertToPincecone,QueryPinecone,
// DeletePinecone
var (
	indexHost = "index Host"
	apiKey    = "api Key"
)

func UpsertToPinecone(item types.Vectorizable) *errors.Error {
	embeddingResult, err := item.Vectorize()
	if err != nil {
		return err
	}

	upsertBody, _ := json.Marshal(map[string]interface{}{
		"vectors": []types.EmbeddingResult{
			embeddingResult,
		},
		"namespace": item.Namespace(),
	})
	requestBody := bytes.NewBuffer(upsertBody)

	req, _err := http.NewRequest("POST", fmt.Sprintf("%s/vectors/upsert", indexHost), requestBody)
	if _err != nil {
		return nil
	}

	req.Header.Set("Api-Key", apiKey)
	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json")

	resp, _err := http.DefaultClient.Do(req)
	if _err != nil {
		return nil
	}

	if resp.StatusCode != 200 {
		return &errors.FailedToUpsertClubToPinecone
	} else {
		return nil
	}
}

func DeleteFromPinecone(item types.Vectorizable) *errors.Error {
	deleteBody, _ := json.Marshal(map[string]interface{}{
		"deleteAll": false,
		"ids": []string{
			item.ID(),
		},
		"namespace": item.Namespace(),
	})

	requestBody := bytes.NewBuffer(deleteBody)

	req, _err := http.NewRequest("POST", fmt.Sprintf("%s/vectors/delete", indexHost), requestBody)
	if _err != nil {
		return nil
	}

	req.Header.Set("Api-Key", apiKey)
	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json")

	resp, _err := http.DefaultClient.Do(req)
	if _err != nil {
		return nil
	}

	if resp.StatusCode != 200 {
		return nil
	} else {
		return nil
	}
}

func SearchPinecone(item types.Vectorizable, topKResults int) ([]string, *errors.Error) {
	searchBody, _ := json.Marshal(map[string]interface{}{
		"includeValues":   false,
		"includeMetadata": false,
		"topK":            topKResults,
		"ids": []string{
			item.ID(),
		},
		"namespace": item.Namespace(),
	})

	requestBody := bytes.NewBuffer(searchBody)

	req, _err := http.NewRequest("POST", fmt.Sprintf("%s/vectors/query", indexHost), requestBody)
	if _err != nil {
		return []string{}, nil
	}

	req.Header.Set("Api-Key", apiKey)
	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json")

	resp, _err := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	if _err != nil {
		return []string{}, nil
	}

	if resp.StatusCode != 200 {
		return []string{}, nil
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
	_err = json.NewDecoder(resp.Body).Decode(&results)
	if _err != nil {
		return []string{}, nil
	}

	resultsToReturn := []string{}
	for i := 0; i < len(results.Matches); i += 1 {
		resultsToReturn = append(resultsToReturn, results.Matches[i].Id)
	}

	return resultsToReturn, nil
}
