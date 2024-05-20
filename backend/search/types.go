// Types that are needed to interact with OpenSearch's API.
package search

import (
	"strings"
	"time"

	"github.com/GenerateNU/sac/backend/entities/models"
)

type Json map[string]interface{}

// Used for making OpenSearch GET /<index>/_search requests.
// OpenSearch's DSL is such that we can't make this a nice static struct with all the fields
// specified; depending on what we are looking to search we'll need to add different clauses.
type SearchEndpointRequest Json

// Used for SearchEndpointRequests.
type BooleanQuery struct {
	Must   []Json `json:"must"`
	Should []Json `json:"should"`
	Filter []Json `json:"filter"`
}

// Used for responses from OpenSearch GET /<index>/_search requests.
type SearchEndpointResponse struct {
	Hits struct {
		Hits []struct {
			Index string  `json:"_index"`
			Score float32 `json:"_score"`
			Id    string  `json:"_id"`
		} `json:"hits"`
	} `json:"hits"`
}

// How Clubs are represented in the OpenSearch index.
type ClubSearchDocument struct {
	Name        string   `json:"name"`
	Preview     string   `json:"preview"`
	Description string   `json:"description"`
	NumMembers  int      `json:"num_members"`
	Tags        []string `json:"tags"`
}

func ClubSearchDocumentFromClub(c *models.Club) ClubSearchDocument {
	tagIds := make([]string, len(c.Tag))
	for i, tag := range c.Tag {
		tagIds[i] = tag.ID.String()
	}

	return ClubSearchDocument{
		Name:        c.Name,
		Preview:     c.Preview,
		Description: c.Description,
		NumMembers:  c.NumMembers,
		Tags:        tagIds,
	}
}

// Query parameters for an API GET /api/v1/search/clubs request.
type ClubSearchQuery struct {
	Search     string   `query:"search"`
	MinMembers int      `query:"min_members"`
	MaxMembers int      `query:"max_members"`
	Tags       []string `query:"tags"`
}

// To Query DSL JSON
func (q *ClubSearchQuery) ToSearchEndpointRequest() SearchEndpointRequest {
	query := BooleanQuery{}

	if q.Search != "" {
		query.Must = append(query.Must, Json{
			"query_string": Json{
				"query": q.Search,
			},
		})
	}

	maxMembers := 0
	if q.MaxMembers != 0 {
		maxMembers = q.MaxMembers
	} else {
		maxMembers = 16384 // FIXME: may break in the future (will a club ever have more than 16384 members?)
	}
	query.Filter = append(query.Filter, Json{
		"range": Json{
			"num_members": Json{
				"lte": maxMembers,
				// by default this 0, hence why we dont check it similar to maxMembers
				"gte": q.MinMembers,
			},
		},
	})

	if len(q.Tags) != 0 {
		query.Filter = append(query.Filter, Json{
			"match": Json{
				"tags": Json{
					"query":                strings.Join(q.Tags, " "),
					"operator":             "or",
					"minimum_should_match": 1,
				},
			},
		})
	}

	return SearchEndpointRequest{
		"query": Json{
			"bool": query,
		},
	}
}

// The (successful) response of an API GET /api/v1/search/clubs request.
type ClubSearchResult struct {
	Results []models.Club `json:"results"`
}

// How Events are represented in the OpenSearch index.
type EventSearchDocument struct {
	Name        string           `json:"name"`
	Preview     string           `json:"preview"`
	Description string           `json:"description"`
	EventType   models.EventType `json:"event_type"`
	StartTime   time.Time        `json:"start_time"`
	EndTime     time.Time        `json:"end_time"`
	Clubs       []string         `json:"clubs"`
	Tags        []string         `json:"tags"`
}

func EventSearchDocumentFromEvent(c *models.Event) EventSearchDocument {
	tagIds := make([]string, len(c.Tag))
	for i, tag := range c.Tag {
		tagIds[i] = tag.ID.String()
	}

	clubIds := make([]string, len(c.Clubs))
	for i, club := range c.Clubs {
		clubIds[i] = club.ID.String()
	}

	return EventSearchDocument{
		Name:        c.Name,
		Preview:     c.Preview,
		Description: c.Description,
		EventType:   c.EventType,
		StartTime:   c.StartTime,
		EndTime:     c.EndTime,
		Tags:        tagIds,
		Clubs:       clubIds,
	}
}

// Query parameters for an API GET /api/v1/search/events request.
type EventSearchQuery struct {
	Search    string             `query:"search"`
	StartTime time.Time          `query:"start_time"`
	EndTime   time.Time          `query:"end_time"`
	EventType []models.EventType `query:"event_type"`
	Clubs     []string           `query:"clubs"`
	Tags      []string           `query:"tags"`
}

// To Query DSL JSON
func (q *EventSearchQuery) ToSearchEndpointRequest() SearchEndpointRequest {
	query := BooleanQuery{}

	if q.Search != "" {
		query.Must = append(query.Must, Json{
			"query_string": Json{
				"query": q.Search,
			},
		})
	}

	if !q.StartTime.IsZero() {
		query.Filter = append(query.Filter, Json{
			"range": Json{
				"start_time": Json{
					"gte": q.StartTime,
				},
			},
		})
	}

	if !q.EndTime.IsZero() {
		query.Filter = append(query.Filter, Json{
			"range": Json{
				"end_time": Json{
					"lte": q.EndTime,
				},
			},
		})
	}

	if len(q.EventType) != 0 {
		query.Filter = append(query.Filter, Json{
			"terms": Json{
				"event_type": q.EventType,
			},
		})
	}

	if len(q.Clubs) != 0 {
		query.Filter = append(query.Filter, Json{
			"match": Json{
				"tags": Json{
					"query":                strings.Join(q.Clubs, " "),
					"operator":             "or",
					"minimum_should_match": 1,
				},
			},
		})
	}

	if len(q.Tags) != 0 {
		query.Filter = append(query.Filter, Json{
			"match": Json{
				"tags": Json{
					"query":                strings.Join(q.Tags, " "),
					"operator":             "or",
					"minimum_should_match": 1,
				},
			},
		})
	}

	return SearchEndpointRequest{
		"query": Json{
			"bool": query,
		},
	}
}

// The (successful) response of an API GET /api/v1/search/events request.
type EventSearchResult struct {
	Results []models.Event `json:"results"`
}

// Used for making OpenSearch /_bulk requests.
type BulkRequestCreate struct {
	Create struct {
		Index string `json:"_index"`
		Id    string `json:"_id"`
	} `json:"create"`
}
