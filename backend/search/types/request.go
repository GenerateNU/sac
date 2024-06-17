package types

import (
	"time"

	"github.com/GenerateNU/sac/backend/entities/models"
)

// JSON parameters for an API GET /api/v1/search/clubs request.
type ClubSearchRequest struct {
	Search     string   `json:"search"`
	MinMembers int      `json:"min_members"`
	MaxMembers int      `json:"max_members"`
	Tags       []string `json:"tags"`
}

type EventSearchRequest struct {
	Search    string             `json:"search"`
	StartTime time.Time          `json:"start_time"`
	EndTime   time.Time          `json:"end_time"`
	EventType []models.EventType `json:"event_type"`
	Clubs     []string           `json:"clubs"`
	Tags      []string           `json:"tags"`
}
