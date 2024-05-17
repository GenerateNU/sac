// Types for incoming requests and outgoing responses.
package search

import (
	"github.com/GenerateNU/sac/backend/entities/models"
)

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
func (q *ClubSearchQuery) ToRequest() SearchEndpointRequest {
	var req SearchEndpointRequest

	req.Query.QueryString.Query = q.Search

	return req
}

type ClubSearchResult struct {
	Results []models.Club `json:"results"`
}
