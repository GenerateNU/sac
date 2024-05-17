package base

import (
	"encoding/json"
	"fmt"

	"github.com/GenerateNU/sac/backend/constants"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/search"
	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
)

func SearchClubs(db *gorm.DB, query search.ClubSearchQuery) (*search.ClubSearchResult, error) {
	_, err := query.ToJson()

	if err != nil {
		return nil, err
	}

	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody("xasdasfaf").
		Get(fmt.Sprintf("%s/clubs/_search", constants.SEARCH_URL))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		// TODO: tf i put here for error
		return nil, nil
	}

	var responseData search.SearchEndpointResponse

	err = json.Unmarshal(resp.Body(), &responseData)
	if err != nil {
		return nil, err
	}

	ids := make([]string, len(responseData.Hits.Hits))
	for i, result := range responseData.Hits.Hits {
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
