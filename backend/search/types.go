// Types that are needed to interact with OpenSearch's API.
package search

import (
	"github.com/GenerateNU/sac/backend/entities/models"
)

// The (successful) response of an API GET /api/v1/search/events request.
type EventSearchResult struct {
	Results []models.Event `json:"results"`
}
