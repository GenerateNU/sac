// Types that are needed to interact with OpenSearch's API.
package search

import (
	"strings"

	"github.com/GenerateNU/sac/backend/entities/models"
)

type Json map[string]interface{}

// Used for making OpenSearch GET /<index>/_search requests.
// OpenSearch's DSL is such that we can't make this a nice static struct with all the fields
// specified; depending on what we are looking to search we'll need to add different clauses.
type SearchEndpointRequest Json

// Used in SearchEndpointRequest
type TagFilterInSearchEndpointRequest struct {
	Term struct {
		Tag string `json:"tag"`
	} `json:"term"`
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
	type BooleanQuery struct {
		Must   []Json `json:"must"`
		Should []Json `json:"should"`
		Filter []Json `json:"filter"`
	}

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

// Used for making OpenSearch /_bulk requests.
type BulkRequestCreate struct {
	Create struct {
		Index string `json:"_index"`
		Id    string `json:"_id"`
	} `json:"create"`
}
