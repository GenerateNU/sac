package base

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"

	json "github.com/goccy/go-json"

	"github.com/GenerateNU/sac/backend/constants"
	"github.com/GenerateNU/sac/backend/search/types"
	"github.com/GenerateNU/sac/backend/utilities"
	"gorm.io/gorm"
)

// Do a GET request to the OpenSearch instance.
func doSearchGetRequest[T, V any](url string, requestBody T) (*V, error) {
	payload, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	resp, err := utilities.Request(http.MethodGet, fmt.Sprintf("%s%s", constants.SEARCH_URI, url), payload, utilities.JSON())
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

func Search[T types.Searchable](db *gorm.DB, query types.SearchRequest) (*types.SearchResult[T], error) {
	result, err := doSearchGetRequest[types.SearchEndpointRequest, types.SearchEndpointResponse](fmt.Sprintf("/%s/_search", query.Index()), query.ToSearchEndpointRequest())
	if err != nil {
		return nil, nil
	}

	ids := make([]string, len(result.Hits.Hits))
	for i, result := range result.Hits.Hits {
		ids[i] = result.Id
	}

	var results []T

	if err = query.Preload(db).Where("id IN ?", ids).Find(&results).Error; err != nil {
		return nil, nil
	}

	return &types.SearchResult[T]{
		Results: results,
	}, nil
}

func Upsert[T types.Searchable](db *gorm.DB, index string, uuid string, model types.ToSearchDocument) error {
	var elem T

	if err := db.Model(&model).Preload("Tag").Where("Clubs").Where("id = ?", uuid).Find(&elem).Error; err != nil {
		return err
	}

	doc := model.ToSearchDocument()
	requestBody := types.Json{
		"doc":           doc,
		"doc_as_upsert": true,
	}
	payload, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	resp, err := utilities.Request(http.MethodPost, fmt.Sprintf("%s/%s/_update/%s", constants.SEARCH_URI, index, uuid), payload, utilities.JSON())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func Delete(db *gorm.DB, index string, uuid string) error {
	resp, err := utilities.Request(http.MethodDelete, fmt.Sprintf("%s/%s/_doc/%s", constants.SEARCH_URI, index, uuid), nil, utilities.JSON())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	slog.Info("delete successful", "response", string(responseBody), "index", index, "uuid", uuid)

	return nil
}
