package types

import (
	"strings"
	"time"

	"github.com/GenerateNU/sac/backend/constants"
	"github.com/GenerateNU/sac/backend/entities/models"
	"gorm.io/gorm"
)

type SearchRequest interface {
	ToSearchEndpointRequest() SearchEndpointRequest
	Index() string
	Preload(db *gorm.DB) *gorm.DB
}

// Used for making OpenSearch GET /<index>/_search requests.
// OpenSearch's DSL is such that we can't make this a nice static struct with all the fields
// specified; depending on what we are looking to search we'll need to add different clauses.
type SearchEndpointRequest Json

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

// JSON parameters for an API GET /api/v1/search/clubs request.
type ClubSearchRequest struct {
	Search     string   `json:"search"`
	MinMembers int      `json:"min_members"`
	MaxMembers int      `json:"max_members"`
	Tags       []string `json:"tags"`
}

// To Query DSL JSON
func (q *ClubSearchRequest) ToSearchEndpointRequest() SearchEndpointRequest {
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
		maxMembers = constants.SEARCH_QUERY_DEFAULT_MAX_MEMBERS // FIXME: may break in the future (will a club ever have more than 16384 members?)
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

func (q *ClubSearchRequest) Preload(db *gorm.DB) *gorm.DB {
	return db.Model(&models.Club{}).Preload("Tag")
}

func (q *ClubSearchRequest) Index() string {
	return "clubs"
}

// JSON parameters for an API GET /api/v1/search/events request.
type EventSearchRequest struct {
	Search    string             `json:"search"`
	StartTime time.Time          `json:"start_time"`
	EndTime   time.Time          `json:"end_time"`
	EventType []models.EventType `json:"event_type"`
	Clubs     []string           `json:"clubs"`
	Tags      []string           `json:"tags"`
}

// To Query DSL JSON
func (q *EventSearchRequest) ToSearchEndpointRequest() SearchEndpointRequest {
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

func (q *EventSearchRequest) Preload(db *gorm.DB) *gorm.DB {
	return db.Model(&models.Event{}).Preload("Tag").Preload("Clubs")
}

func (q *EventSearchRequest) Index() string {
	return "events"
}
