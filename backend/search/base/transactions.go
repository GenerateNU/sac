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

// Do a GET request to the OpenSearch instance.
func doSearchGetRequest[T, V any](url string, requestBody T) (*V, error) {
	payload, err := json.Marshal(requestBody)
	print(string(payload))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", constants.SEARCH_URL, url), bytes.NewBuffer(payload))
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

func SearchClubs(db *gorm.DB, query search.ClubSearchRequest) (*search.ClubSearchResult, error) {
	result, err := doSearchGetRequest[search.SearchEndpointRequest, search.SearchEndpointResponse]("/clubs/_search", query.ToSearchEndpointRequest())
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

func SearchEvents(db *gorm.DB, query search.EventSearchRequest) (*search.EventSearchResult, error) {
	result, err := doSearchGetRequest[search.SearchEndpointRequest, search.SearchEndpointResponse]("/events/_search", query.ToSearchEndpointRequest())
	if err != nil {
		return nil, nil
	}

	ids := make([]string, len(result.Hits.Hits))
	for i, result := range result.Hits.Hits {
		ids[i] = result.Id
	}

	var events []models.Event

	if err = db.Model(&models.Event{}).Preload("Tag").Where("id IN ?", ids).Find(&events).Error; err != nil {
		return nil, nil
	}

	return &search.EventSearchResult{
		Results: events,
	}, nil
}

func UpsertClubToIndex(db *gorm.DB, clubUuid string) error {
	var club models.Club

	if err := db.Where("id = ?", clubUuid).Find(&club).Error; err != nil {
		return err
	}

	doc := search.ClubSearchDocumentFromClub(&club)
	requestBody := search.Json{
		"doc":           doc,
		"doc_as_upsert": true,
	}
	payload, err := json.Marshal(requestBody)

	print(string(payload))
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s/_update/%s", constants.SEARCH_URL, "clubs", club.ID), bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	print(string(responseBody))

	return nil
}

func DeleteClubFromIndex(db *gorm.DB, clubUuid string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s/_doc/%s", constants.SEARCH_URL, "clubs", clubUuid), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	print(string(responseBody))

	return nil
}

func UpsertEventToIndex(db *gorm.DB, eventUuid string) error {
	var event models.Event

	if err := db.Where("id = ?", eventUuid).Find(&event).Error; err != nil {
		return err
	}

	doc := search.EventSearchDocumentFromEvent(&event)
	requestBody := search.Json{
		"doc":           doc,
		"doc_as_upsert": true,
	}
	payload, err := json.Marshal(requestBody)

	print(string(payload))
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s/_update/%s", constants.SEARCH_URL, "events", event.ID), bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	print(string(responseBody))

	return nil
}

func DeleteEventFromIndex(db *gorm.DB, eventUuid string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s/_doc/%s", constants.SEARCH_URL, "events", eventUuid), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	print(string(responseBody))

	return nil
}
