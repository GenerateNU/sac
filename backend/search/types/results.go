package types

import "github.com/GenerateNU/sac/backend/entities/models"

type SearchResult[T models.Event | models.Club] struct {
	Results []T `json:"results"`
}
