// Types that are needed to interact with OpenSearch's API.
package search

import "github.com/GenerateNU/sac/backend/entities/models"

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
	var req SearchEndpointRequest

	/*req.Query.Bool.Must.QueryString.Query = q.Search

	req.Query.Bool.Filter.Range.NumMembers.MinMembers = q.MinMembers

	if q.MaxMembers != 0 {
		req.Query.Bool.Filter.Range.NumMembers.MaxMembers = q.MaxMembers
	} else {
		// FIXME: This is probably going to break in the future. We should probably be figuring out
		// what the most members of any club is and set this to the default.
		req.Query.Bool.Filter.Range.NumMembers.MaxMembers = 16384
	}

	for _, tag := range q.Tags {
		req.Query.Bool.Must.Terms.Tags = append(req.Query.Bool.Must.Terms.Tags, tag)
	}

	req.Query.Bool.Must.Terms.MinimumShouldMatch = 1*/
	// TODO:
	// must - query_string, filter - tags (done in a match) + members (done in a range)
	return req
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
