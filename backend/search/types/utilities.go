package types

import "github.com/GenerateNU/sac/backend/entities/models"

type Json map[string]interface{}

// Used for SearchEndpointRequests.
type BooleanQuery struct {
	Must   []Json `json:"must"`
	Should []Json `json:"should"`
	Filter []Json `json:"filter"`
}

// Used for making OpenSearch /_bulk requests.
type BulkRequestCreate struct {
	Create struct {
		Index string `json:"_index"`
		Id    string `json:"_id"`
	} `json:"create"`
}

type Searchable interface {
	models.Club | models.Event
}
