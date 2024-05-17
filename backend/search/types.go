// Types that are needed to interact with OpenSearch's API.
package search

import "github.com/GenerateNU/sac/backend/entities/models"

// Used for making OpenSearch GET /<index>/_search requests.
type SearchEndpointRequest struct {
	Query struct {
		QueryString struct {
			Query string `json:"query"`
		} `json:"query_string"`
	} `json:"query"`
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
	Search string `query:"search"`
	// Tags                    []string          `query:"tags"`
	// MinMembers              int               `query:"min_members"`
	// MaxMembers              int               `query:"max_members"`
	// RecruitmentCycle        *RecruitmentCycle `query:"recruitment_cycle"`
	// IsRecruiting            *bool             `query:"is_recruiting"`
	// MinWeeklyTimeCommitment int               `query:"min_weekly_time_commitment"`
	// MaxWeeklyTimeCommitment int               `query:"max_weekly_time_commitment"`
	// OneWordToDescribeUs     string            `query:"one_word_to_describe_us"`
	// Search                  string            `query:"search"`
}

// To Query DSL JSON
func (q *ClubSearchQuery) ToSearchEndpointRequest() SearchEndpointRequest {
	var req SearchEndpointRequest

	req.Query.QueryString.Query = q.Search

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
