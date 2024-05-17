package base

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/GenerateNU/sac/backend/constants"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/search"
	"gorm.io/gorm"
)

// Do a request to the OpenSearch instance.
func doOpenSearchHttpRequest[T, V any](url string, requestBody T) (*V, error) {
	payload, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/clubs/_search", constants.SEARCH_URL), bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var responseData V
	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func SearchClubs(db *gorm.DB, query search.ClubSearchQuery) (*search.ClubSearchResult, error) {
	result, err := doOpenSearchHttpRequest[search.SearchEndpointRequest, search.SearchEndpointResponse]("/clubs/_search", query.ToRequest())
	if err != nil {
		return nil, nil
	}

	ids := make([]string, len(result.Hits.Hits))
	for i, result := range result.Hits.Hits {
		ids[i] = result.Id
	}

	var clubs []models.Club

	if err = db.Model(&models.Club{}).Preload("Tag").Where("id IN ?", ids).Find(&clubs).Error; err != nil {
		return nil, nil
	}

	return &search.ClubSearchResult{
		Results: clubs,
	}, nil
}
